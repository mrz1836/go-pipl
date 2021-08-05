// Package pipl provides a way to interact programmatically with the Pipl API in Golang.
// For more detailed information on the Pipl search API and what we're actually
// wrapping, check out their official API reference: https://docs.pipl.com/reference/#overview
package pipl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SourceLevel is used internally to represent the possible values
// for show_sources in queries to be submitted: {"all", "matching"/"true", "false"}
type SourceLevel string

// MatchRequirements specifies the conditions for a successful person match in our search.
// This is useful for saving money with the Pipl API, as you only need to pay for the
// data you wanted back. If your search results didn't satisfy the match requirements, then
// no data is returned, and you don't pay.
type MatchRequirements string

// SourceCategoryRequirements specifies the data categories that must be included in
// results for a successful match. If there is no data from the requested categories,
// then the results returned are empty, and you're not charged.
type SourceCategoryRequirements string

// SearchParameters holds options that can affect data returned by a search.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#configuration-parameters
type SearchParameters struct {
	// apiKey is required
	apiKey string

	// ShowSources specifies the level of sources info to return with search results, one of {ShowSourcesMatching, ShowSourcesAll, ShowSourcesNone}
	ShowSources SourceLevel

	// MatchRequirements specifies the criteria for a successful Person match.
	// Results that don't fit your match requirements are discarded. If the remaining
	// search results would be empty, you are not charged for the query.
	MatchRequirements MatchRequirements

	// SourceCategoryRequirements specifies the data categories that must be included in
	// results for a successful match. If there is no data from the requested categories,
	// then the results returned are empty, and you're not charged.
	SourceCategoryRequirements SourceCategoryRequirements

	// MinimumProbability is the minimum acceptable probability for inferred data
	MinimumProbability float32

	// MinimumMatch specifies the minimum match confidence for a possible person to be returned in search results
	MinimumMatch float32

	// InferPersons specifies whether the Pipl should return results inferred by statistical analysis
	InferPersons bool

	// HideSponsored specifies whether to omit sponsored data from search results
	HideSponsored bool

	// LiveFeeds specifies whether to use live data sources
	LiveFeeds bool

	// Returns the best high ranking match to your search. API will return either a Person (when high scoring profile is found) or a No Match
	TopMatch bool
}

// ThumbnailSettings is for the thumbnail url settings to be automatically returned
// if any images are found and meet the criteria
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Example: http://thumb.pipl.com/image?height=250&width=250&favicon=true&zoom_face=true&tokens=FIRST_TOKEN,SECOND_TOKEN
type ThumbnailSettings struct {

	// URL is the thumbnail url
	URL string

	// Height of the image
	Height int

	// Width of the image
	Width int

	// Enabled (detects images, automatically adds thumbnail urls)
	Enabled bool

	// Favicon if the icon should be shown or not
	Favicon bool

	// ZoomFace is whether to enable face zoom.
	ZoomFace bool
}

// NewClient creates a new search client to submit queries with
// Parameters values are set to the defaults defined by Pipl.
//
// For more information: https://docs.pipl.com/reference#configuration-parameters
func NewClient(apiKey string, clientOptions *Options) (c *Client, err error) {

	// Test for the key
	if len(apiKey) == 0 {
		err = fmt.Errorf("api key must be set, %s", "api_key")
		return
	}

	// Create a client using the given options
	c = createClient(clientOptions)

	// Create default search parameters
	c.Parameters.Search = new(SearchParameters)
	c.Parameters.Search.apiKey = apiKey
	c.Parameters.Search.HideSponsored = true
	c.Parameters.Search.InferPersons = false
	c.Parameters.Search.TopMatch = false
	c.Parameters.Search.LiveFeeds = true
	c.Parameters.Search.MatchRequirements = MatchRequirementsNone
	c.Parameters.Search.MinimumMatch = MinimumMatch
	c.Parameters.Search.MinimumProbability = MinimumProbability
	c.Parameters.Search.ShowSources = ShowSourcesAll // ShowSourcesNone
	c.Parameters.Search.SourceCategoryRequirements = SourceCategoryRequirementsNone

	// Create default thumbnail parameters (thumbnail url functionality)
	c.Parameters.Thumbnail = new(ThumbnailSettings)
	c.Parameters.Thumbnail.Enabled = false
	c.Parameters.Thumbnail.Height = ThumbnailHeight
	c.Parameters.Thumbnail.URL = thumbnailEndpoint
	c.Parameters.Thumbnail.Width = ThumbnailWidth

	// Return the client
	return
}

// SearchMeetsMinimumCriteria is used internally by Search to do some very
// basic verification that the verify that search object has enough terms to
// meet the requirements for a search.
// From Pipl documentation:
// 		"The minimal requirement to run a search is to have at least one full
//		name, email, phone, username, user_id, URL or a single valid US address
//		(down to a house number). We can’t search for a job title or location
//		alone. We’re not a directory and can't provide bulk lists of people,
//		rather we specialize in identity resolution of single individuals."
func SearchMeetsMinimumCriteria(searchPerson *Person) bool {

	// If an email is found, that meets minimum criteria
	if searchPerson.HasEmail() {
		return true
	}

	// If a phone is found, that meets minimum criteria
	if searchPerson.HasPhone() {
		return true
	}

	// If a userID is found, that meets minimum criteria
	if searchPerson.HasUserID() {
		return true
	}

	// If a username is found, that meets minimum criteria
	if searchPerson.HasUsername() {
		return true
	}

	// If a URL is found, that meets minimum criteria
	if searchPerson.HasURL() {
		return true
	}

	// If a full name is found, that meets minimum criteria
	if searchPerson.HasName() {
		return true
	}

	// If an address is found, that meets minimum criteria
	if searchPerson.HasAddress() {
		return true
	}

	// Did not meet criteria, fail
	return false
}

// Search takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If successful, the response struct
// will contain the results, and err will be nil. If an error occurs, the struct pointer
// will be nil, and you should check err for additional information. This method will only
// return one full person, and a preview of possible people if < 100% match. Use the SearchAllPossiblePeople()
// method to get all the details when searching.
func (c *Client) Search(searchPerson *Person) (response *Response, err error) {

	// Do we meet the minimum requirements for searching?
	if !SearchMeetsMinimumCriteria(searchPerson) {
		err = fmt.Errorf("the search request submitted does not contain enough sufficient terms. " +
			"You must have one of the following: full name, email, phone, username, userID, url, or full street address")
		return
	}

	// Start the post data
	postData := url.Values{}

	// Add the API key
	postData.Add("key", c.Parameters.Search.apiKey)

	// Do not return formatted
	postData.Add("pretty", "false")

	// Should we show sources?
	if c.Parameters.Search.ShowSources != ShowSourcesNone {
		postData.Add("show_sources", string(c.Parameters.Search.ShowSources))
	}

	// Add match requirements?
	if c.Parameters.Search.MatchRequirements != MatchRequirementsNone {
		postData.Add("match_requirements", string(c.Parameters.Search.MatchRequirements))
	}

	// Add source category requirements?
	if c.Parameters.Search.SourceCategoryRequirements != SourceCategoryRequirementsNone {
		postData.Add("source_category_requirements", string(c.Parameters.Search.SourceCategoryRequirements))
	}

	// Ask for the top match?
	if c.Parameters.Search.TopMatch {
		postData.Add("top_match", "true")
	}

	// Parse the search object
	var personJSON []byte
	if personJSON, err = json.Marshal(searchPerson); err != nil {
		return
	}

	// Add the person to the request
	postData.Add("person", string(personJSON))

	// Fire the request
	return c.PiplRequest(searchAPIEndpoint, http.MethodPost, &postData)
}

// SearchAllPossiblePeople takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If possible people are found, they are also
// looked up using the SearchByPointer()
func (c *Client) SearchAllPossiblePeople(searchPerson *Person) (response *Response, err error) {

	// Lookup the person(s)
	if response, err = c.Search(searchPerson); err != nil {
		return
	}

	// When multiple PossiblePersons are returned, we get a "preview" of each of them (< 100% match confidence)
	if response.PersonsCount > 1 {
		for index, person := range response.PossiblePersons {

			// In order to get the full info on each, we need to a follow-up query
			// to pull a full person profile by search pointer
			searchPointer := person.SearchPointer
			var searchResponse *Response
			if searchResponse, err = c.SearchByPointer(searchPointer); err != nil {
				return
			}

			// Replace the preview with the full details
			response.PossiblePersons[index] = searchResponse.Person
		}
	}

	return
}

// SearchByPointer takes a search pointer string and returns the full
// information for the person associated with that pointer
func (c *Client) SearchByPointer(searchPointer string) (response *Response, err error) {

	// So we have a search pointer?
	if len(searchPointer) < 20 {
		err = fmt.Errorf("invalid search pointer: %s", searchPointer)
		return
	}

	// Start the post data
	postData := url.Values{}

	// Add the API key
	postData.Add("key", c.Parameters.Search.apiKey)

	// Add the search pointer
	postData.Add("search_pointer", searchPointer)

	// Fire the request
	return c.PiplRequest(searchAPIEndpoint, http.MethodPost, &postData)
}

// PiplRequest is a generic pipl request wrapper that can be used without the constraints
// of the Search or SearchByPointer methods
func (c *Client) PiplRequest(endpoint string, method string, params *url.Values) (response *Response, err error) {

	// Set reader
	var bodyReader io.Reader

	// Switch on method
	switch method {
	case http.MethodPost:
		{
			encodedParams := params.Encode()
			bodyReader = strings.NewReader(encodedParams)
			c.LastRequest.PostData = encodedParams
		}
	case http.MethodGet:
		if params != nil {
			endpoint += "?" + params.Encode()
		}
	}

	// Store for debugging purposes
	c.LastRequest.Method = method
	c.LastRequest.URL = endpoint

	// Start the request
	var request *http.Request
	if request, err = http.NewRequestWithContext(context.Background(), method, endpoint, bodyReader); err != nil {
		return
	}

	// Set the headers
	request.Header.Set("User-Agent", c.Parameters.UserAgent)

	// Set the content type on method
	if method == http.MethodPost {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Fire the http request
	var resp *http.Response
	if resp, err = c.httpClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		_ = resp.Body.Close()
	}()

	// Read the body
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// Parse the response
	response = new(Response)
	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// Thumbnail generation enabled?
	if c.Parameters.Thumbnail.Enabled {

		// Process the current person
		response.Person.ProcessThumbnails(c)

		// Do we have possible persons?
		if len(response.PossiblePersons) > 0 {
			for index := range response.PossiblePersons {
				response.PossiblePersons[index].ProcessThumbnails(c)
			}
		}
	}

	// Done
	return
}
