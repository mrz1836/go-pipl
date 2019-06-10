// Package pipl provides a way to interact programmatically with the Pipl API in Golang.
// For more detailed information on the Pipl search API and what we're actually
// wrapping, check out their official API reference: https://docs.pipl.com/reference/#overview
package pipl

import (
	"encoding/json"
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
// no data is returned and you don't pay.
type MatchRequirements string

// SourceCategoryRequirements specifies the data categories that must be included in
// results for a successful match. If there is no data from the requested categories,
// then the results returned are empty and you're not charged.
type SourceCategoryRequirements string

// Package global constants
const (
	// SearchAPIEndpoint is where we POST queries to
	SearchAPIEndpoint string = "https://api.pipl.com/search/"

	// ShowSourcesNone specifies that we don't need source info back with search results
	ShowSourcesNone SourceLevel = "false"

	// ShowSourcesAll specifies that we want all source info back with our search results
	ShowSourcesAll SourceLevel = "all"

	// ShowSourcesMatching specifies that we want source info that corresponds to data that satisfies our match requirements
	ShowSourcesMatching SourceLevel = "true"

	// MatchRequirementsNone specifies that we don't have any match requirements for this search
	MatchRequirementsNone MatchRequirements = ""

	// MatchRequirementsEmail specifies that we want to match on this field
	MatchRequirementsEmail MatchRequirements = "email"

	// MatchRequirementsPhone specifies that we want to match on this field
	MatchRequirementsPhone MatchRequirements = "phone"

	// MatchRequirementsEmailAndPhone specifies that we want to match on this field
	MatchRequirementsEmailAndPhone MatchRequirements = "email and phone"

	// MatchRequirementsEmailAndName specifies that we want to match on this field
	MatchRequirementsEmailAndName MatchRequirements = "email and name"

	// MatchRequirementsEmailOrPhone specifies that we want to match on this field
	MatchRequirementsEmailOrPhone MatchRequirements = "email or phone"

	// todo: finish adding match criteria - also make this flexible and easier to use
	// https://docs.pipl.com/reference#match-criteria

	// MinimumProbability is the score for probability
	MinimumProbability = 0.9

	// MinimumMatch is the minimum for a match
	MinimumMatch = 0.0

	// SourceCategoryRequirementsNone specifies that we don't require any specific sources in our results.
	SourceCategoryRequirementsNone SourceCategoryRequirements = ""

	// SourceCategoryRequirementsProfessionalAndBusiness is used for: match_requirements=(emails and jobs)
	SourceCategoryRequirementsProfessionalAndBusiness SourceCategoryRequirements = "professional_and_business"
)

// Client holds client configuration settings
type Client struct {
	// HTTPClient carries out the POST operations
	HTTPClient *http.Client
	// Parameters contains the search parameters that are submitted with your query,
	// which may affect the data returned
	SearchParameters *SearchParameters
}

// SearchParameters holds options that can affect data returned by a search.
type SearchParameters struct {
	// APIKey is required
	APIKey string

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

// NewClient creates a new search client to submit queries with.
// Parameters values are set to the defaults defined by Pipl.
// For more information:
// https://docs.pipl.com/reference#configuration-parameters
func NewClient(APIKey string) (client *Client) {
	client = new(Client)
	client.HTTPClient = new(http.Client)
	client.SearchParameters = new(SearchParameters)
	client.SearchParameters.APIKey = APIKey
	//client.SearchParameters.HideSponsored = false
	client.SearchParameters.HideSponsored = true
	client.SearchParameters.InferPersons = false
	client.SearchParameters.LiveFeeds = true
	client.SearchParameters.MatchRequirements = MatchRequirementsNone
	client.SearchParameters.MinimumMatch = MinimumMatch
	client.SearchParameters.MinimumProbability = MinimumProbability
	client.SearchParameters.ShowSources = ShowSourcesAll
	//client.SearchParameters.ShowSources = ShowSourcesNone
	client.SearchParameters.SourceCategoryRequirements = SourceCategoryRequirementsNone
	return
}

// SearchMeetsMinimumCriteria is used internally by SearchByPerson to do some very
// basic verification that the verify that search object has enough terms to
// meet the requirements for a search.
// From Pipl documentation:
// 		"The minimal requirement to run a search is to have at least one full
//		name, email, phone, username, user_id, URL or a single valid US address
//		(down to a house number). We can’t search for a job title or location
//		alone. We’re not a directory and can't provide bulk lists of people,
//		rather we specialize in identity resolution of single individuals."
func SearchMeetsMinimumCriteria(p *Person) bool {

	// If an email is found, that meets minimum criteria
	if p.HasEmail() {
		return true
	}

	// If a phone is found, that meets minimum criteria
	if p.HasPhone() {
		return true
	}

	// If a userID is found, that meets minimum criteria
	if p.HasUserID() {
		return true
	}

	// If a username is found, that meets minimum criteria
	if p.HasUsername() {
		return true
	}

	// If a URL is found, that meets minimum criteria
	if p.HasURL() {
		return true
	}

	// If a full name is found, that meets minimum criteria
	if p.HasName() {
		return true
	}

	// If an address is found, that meets minimum criteria
	if p.HasAddress() {
		return true
	}

	// Did not meet criteria, fail
	return false
}

// SearchByPerson takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If successful, the response struct
// will contains the results, and err will be nil. If an error occurs, the struct pointer
// will be nil and you should check err for additional information.
func (client *Client) SearchByPerson(searchObject *Person) (response *Response, err error) {
	if !SearchMeetsMinimumCriteria(searchObject) {
		return nil, &ErrInsufficientSearch{}
	}
	postData := url.Values{}
	postData.Add("key", client.SearchParameters.APIKey)

	if client.SearchParameters.ShowSources != ShowSourcesNone {
		postData.Add("show_sources", string(client.SearchParameters.ShowSources))
	}
	if client.SearchParameters.MatchRequirements != MatchRequirementsNone {
		postData.Add("match_requirements", string(client.SearchParameters.MatchRequirements))
	}
	if client.SearchParameters.SourceCategoryRequirements != SourceCategoryRequirementsNone {
		postData.Add("source_category_requirements", string(client.SearchParameters.SourceCategoryRequirements))
	}
	var personJSON []byte
	personJSON, err = json.Marshal(searchObject)
	if err != nil {
		return
	}
	postData.Add("person", string(personJSON))
	var request *http.Request
	request, err = http.NewRequest("POST", SearchAPIEndpoint, strings.NewReader(postData.Encode()))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	var resp *http.Response
	resp, err = client.HTTPClient.Do(request)
	if err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response = new(Response)
	err = json.Unmarshal(body, response)
	return
}

// SearchByPointer takes a search pointer string and returns the full
// information for the person associated with that pointer
func (client *Client) SearchByPointer(searchPointer string) (person *Person, err error) {
	postData := url.Values{}
	postData.Add("key", client.SearchParameters.APIKey)
	postData.Add("search_pointer", searchPointer)
	var request *http.Request
	request, err = http.NewRequest("POST", SearchAPIEndpoint, strings.NewReader(postData.Encode()))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	var response *http.Response
	response, err = client.HTTPClient.Do(request)
	if err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	piplResponse := new(Response)
	err = json.Unmarshal(body, piplResponse)
	if err != nil {
		return
	}

	// Set the person from the response
	person = &piplResponse.Person
	return
}
