// Package pipl provides a way to interact programmatically with the Pipl API in Golang.
// For more detailed information on the Pipl search API and what we're actually
// wrapping, check out their official API reference: https://docs.pipl.com/reference/#overview
package pipl

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

// SourceLevel is used internally to represent the possible values
// for show_sources in queries to be submitted: {"all", "matching"/"true", "false"}
type SourceLevel string

// MatchRequirements specifies the conditions for a successful person match in our search.
// This is useful for saving money with the Pipl API, as you only need to pay for the
// data you wanted back. If your search results didn't satisfy the match requirements, then
// no data is returned and you don't pay.
type MatchRequirements string

// SourceCategoryRequirements specifies the data categories that must be included in
// results for a successful match. If there is no data from the requested categories,
// then the results returned are empty and you're not charged.
type SourceCategoryRequirements string

// Client holds client configuration settings
type Client struct {
	// HTTPClient carries out the POST operations
	HTTPClient heimdall.Client

	// SearchParameters contains the search parameters that are submitted with your query,
	// which may affect the data returned
	SearchParameters *SearchParameters

	// ThumbnailSettings is for the thumbnail url settings
	ThumbnailSettings *ThumbnailSettings

	// LastRequest is the raw information from the last request
	LastRequest *LastRequest
}

// SearchParameters holds options that can affect data returned by a search.
//
// Source: https://docs.pipl.com/reference#configuration-parameters
type SearchParameters struct {
	// APIKey is required
	APIKey string

	// UserAgent (optional for changing user agents)
	UserAgent string

	// MinimumProbability is the minimum acceptable probability for inferred data
	MinimumProbability float32

	// InferPersons specifies whether or not the Pipl should return results inferred by statistical analysis
	InferPersons bool

	// MinimumMatch specifies the minimum match confidence for a possible person to be returned in search results
	MinimumMatch float32

	// ShowSources specifies the level of sources info to return with search results, one of {ShowSourcesMatching, ShowSourcesAll, ShowSourcesNone}
	ShowSources SourceLevel

	// HideSponsored specifies whether to omit sponsored data from search results
	HideSponsored bool

	// LiveFeeds specifies whether to use live data sources
	LiveFeeds bool

	// MatchRequirements specifies the criteria for a successful Person match.
	// Results that don't fit your match requirements are discarded. If the remaining
	// search results would be empty, you are not charged for the query.
	MatchRequirements MatchRequirements

	// SourceCategoryRequirements specifies the data categories that must be included in
	// results for a successful match. If there is no data from the requested categories,
	// then the results returned are empty and you're not charged.
	SourceCategoryRequirements SourceCategoryRequirements
}

// ThumbnailSettings is for the thumbnail url settings to be automatically returned
// if any images are found and meet the criteria
//
// Example: http://thumb.pipl.com/image?height=250&width=250&favicon=true&zoom_face=true&tokens=FIRST_TOKEN,SECOND_TOKEN
type ThumbnailSettings struct {
	// Enabled (detects images, automatically adds thumbnail urls)
	Enabled bool

	// URL is the thumbnail url
	URL string

	// Height of the image
	Height int

	// Width of the image
	Width int

	// Favicon if the icon should be shown or not
	Favicon bool

	// ZoomFace is whether to enable face zoom.
	ZoomFace bool
}

// LastRequest is used to track what was submitted to pipl on the piplRequest()
type LastRequest struct {
	// Method is either POST or GET
	Method string

	// PostData is the post data submitted if POST request
	PostData string

	// URL is the url used for the request
	URL string
}

// NewClient creates a new search client to submit queries with.
// Parameters values are set to the defaults defined by Pipl.
//
// For more information: https://docs.pipl.com/reference#configuration-parameters
func NewClient(APIKey string) (c *Client, err error) {

	// Test for the key
	if len(APIKey) == 0 {
		err = fmt.Errorf("api key must be set, %s", APIKey)
		return
	}

	// Create a client
	c = new(Client)

	// Create exponential backoff
	backOff := heimdall.NewExponentialBackoff(
		ConnectionInitialTimeout,
		ConnectionMaxTimeout,
		ConnectionExponentFactor,
		ConnectionMaximumJitterInterval,
	)

	// Create the http client
	//c.HTTPClient = new(http.Client) (@mrz this was the original HTTP client)
	c.HTTPClient = httpclient.NewClient(
		httpclient.WithHTTPTimeout(ConnectionWithHTTPTimeout),
		httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
		httpclient.WithRetryCount(ConnectionRetryCount),
		httpclient.WithHTTPClient(&http.Client{
			Transport: ClientDefaultTransport,
		}),
	)

	// Create default search parameters
	c.SearchParameters = new(SearchParameters)
	c.SearchParameters.APIKey = APIKey
	c.SearchParameters.HideSponsored = true
	c.SearchParameters.InferPersons = false
	c.SearchParameters.LiveFeeds = true
	c.SearchParameters.MatchRequirements = MatchRequirementsNone
	c.SearchParameters.MinimumMatch = MinimumMatch
	c.SearchParameters.MinimumProbability = MinimumProbability
	c.SearchParameters.ShowSources = ShowSourcesAll //ShowSourcesNone
	c.SearchParameters.SourceCategoryRequirements = SourceCategoryRequirementsNone
	c.SearchParameters.UserAgent = DefaultUserAgent

	// Create default thumbnail parameters (thumbnail url functionality)
	c.ThumbnailSettings = new(ThumbnailSettings)
	c.ThumbnailSettings.Enabled = false
	c.ThumbnailSettings.Height = ThumbnailHeight
	c.ThumbnailSettings.URL = ThumbnailEndpoint
	c.ThumbnailSettings.Width = ThumbnailWidth

	// Create a last request struct
	c.LastRequest = new(LastRequest)

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
// will contains the results, and err will be nil. If an error occurs, the struct pointer
// will be nil and you should check err for additional information. This method will only
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
	postData.Add("key", c.SearchParameters.APIKey)

	// Do not return formatted
	postData.Add("pretty", "false")

	// Should we show sources?
	if c.SearchParameters.ShowSources != ShowSourcesNone {
		postData.Add("show_sources", string(c.SearchParameters.ShowSources))
	}

	// Add match requirements?
	if c.SearchParameters.MatchRequirements != MatchRequirementsNone {
		postData.Add("match_requirements", string(c.SearchParameters.MatchRequirements))
	}

	// Add source category requirements?
	if c.SearchParameters.SourceCategoryRequirements != SourceCategoryRequirementsNone {
		postData.Add("source_category_requirements", string(c.SearchParameters.SourceCategoryRequirements))
	}

	// Parse the search object
	var personJSON []byte
	if personJSON, err = json.Marshal(searchPerson); err != nil {
		return
	}

	// Add the person to the request
	postData.Add("person", string(personJSON))

	// Fire the request
	return c.PiplRequest(SearchAPIEndpoint, "POST", &postData)
}

// SearchAllPossiblePeople takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If possible people are found, they are also
// looked up using the SearchByPointer()
func (c *Client) SearchAllPossiblePeople(searchPerson *Person) (response *Response, err error) {

	// Lookup the person(s)
	if response, err = c.Search(searchPerson); err != nil {
		return
	}

	// When multiple PossiblePersons are returned, we get a "preview" of each of
	// each of them (< 100% match confidence)
	if response.PersonsCount > 1 {
		for index, person := range response.PossiblePersons {

			// In order to get the full info on each, we need to a follow up query
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
	postData.Add("key", c.SearchParameters.APIKey)

	// Add the search pointer
	postData.Add("search_pointer", searchPointer)

	// Fire the request
	return c.PiplRequest(SearchAPIEndpoint, "POST", &postData)
}

// PiplRequest is a generic pipl request wrapper that can be used without the constraints
// of the Search or SearchByPointer methods
func (c *Client) PiplRequest(endpoint string, method string, params *url.Values) (response *Response, err error) {

	// Set reader
	var bodyReader io.Reader

	// Switch on POST vs GET
	switch method {
	case "POST":
		{
			encodedParams := params.Encode()
			bodyReader = strings.NewReader(encodedParams)
			c.LastRequest.PostData = encodedParams
		}
	case "GET":
		{
			endpoint += "?" + params.Encode()
		}
	}

	// Store for debugging purposes
	c.LastRequest.Method = method
	c.LastRequest.URL = endpoint

	// Start the request
	var request *http.Request
	if request, err = http.NewRequest(method, endpoint, bodyReader); err != nil {
		return
	}

	// Change the header (user agent is in case they block default Go user agents)
	request.Header.Set("User-Agent", c.SearchParameters.UserAgent)

	// Set the content type on POST
	if method == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Fire the http request
	var resp *http.Response
	if resp, err = c.HTTPClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %s", err.Error())
		}
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
	if c.ThumbnailSettings.Enabled {

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
