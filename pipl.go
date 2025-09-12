// Package pipl provides a way to interact programmatically with the PIPL API in Golang.
// For more detailed information on the PIPL search API and what we're actually
// wrapping, check out their official API reference: https://docs.pipl.com/reference/#overview
package pipl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// Search takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If successful, the response struct
// will contain the results, and err will be nil. If an error occurs, the struct pointer
// will be nil, and you should check err for additional information. This method will only
// return one full person, and a preview of possible people if < 100% match. Use the SearchAllPossiblePeople()
// method to get all the details when searching.
func (c *Client) Search(ctx context.Context, searchPerson *Person) (*Response, error) {
	// Do we meet the minimum requirements for searching?
	if !SearchMeetsMinimumCriteria(searchPerson) {
		return nil, ErrDoesNotMeetMinimumCriteria
	}

	// Start the post data
	postData := url.Values{}

	// Add the API key (always - API is required by default)
	postData.Add(fieldAPIKey, c.options.apiKey)

	// Option for pretty response
	if !c.options.searchOptions.Search.Pretty {
		postData.Add(fieldPretty, valueFalse)
	}

	// Should we show sources?
	if c.options.searchOptions.Search.ShowSources != ShowSourcesNone {
		postData.Add(fieldShowSources, string(c.options.searchOptions.Search.ShowSources))
	}

	// Add match requirements?
	if c.options.searchOptions.Search.MatchRequirements != MatchRequirementsNone {
		postData.Add(fieldMatchRequirements, string(c.options.searchOptions.Search.MatchRequirements))
	}

	// Add source category requirements?
	if c.options.searchOptions.Search.SourceCategoryRequirements != SourceCategoryRequirementsNone {
		postData.Add(fieldSourceCategoryRequirements, string(c.options.searchOptions.Search.SourceCategoryRequirements))
	}

	// Custom minimum match
	if c.options.searchOptions.Search.MinimumMatch != MinimumMatch {
		postData.Add(fieldMinimumMatch, fmt.Sprintf("%v", c.options.searchOptions.Search.MinimumMatch))
	}

	// Set the "hide sponsors" flag (default is false)
	if c.options.searchOptions.Search.HideSponsored {
		postData.Add(fieldHideSponsored, valueTrue)
	}

	// Set the "infer persons" flag (default is false)
	if c.options.searchOptions.Search.InferPersons {
		postData.Add(fieldInferPersons, valueTrue)
	}

	// Ask for the top match?
	if c.options.searchOptions.Search.TopMatch {
		postData.Add(fieldTopMatch, valueTrue)
	}

	// Set the live feeds flag (default is true)
	if !c.options.searchOptions.Search.LiveFeeds {
		postData.Add(fieldLiveFeeds, valueFalse)
	}

	// Parse the search object
	personJSON, err := json.Marshal(searchPerson)
	if err != nil { // This should NEVER error out since the struct is being generated
		return nil, err
	}

	// Add the person to the request
	postData.Add(fieldPerson, string(personJSON))

	// Fire the request
	var response *Response
	response, err = httpRequest(ctx, c, searchAPIEndpoint, &postData)
	if err != nil {
		return nil, err
	} else if len(response.Error) > 0 {
		return nil, fmt.Errorf("%w: %s", ErrAPIResponse, response.Error)
	}
	return response, nil
}

// SearchAllPossiblePeople takes a person object (filled with search terms) and returns the
// results in the form of a Response struct. If possible people are found, they are also
// looked up using the SearchByPointer()
func (c *Client) SearchAllPossiblePeople(ctx context.Context, searchPerson *Person) (response *Response, err error) {
	// Lookup the person(s)
	if response, err = c.Search(ctx, searchPerson); err != nil {
		return
	}

	// When multiple PossiblePersons are returned, we get a "preview" of each of them (< 100% match confidence)
	if response.PersonsCount > 1 {
		for index, person := range response.PossiblePersons {

			// In order to get the full info on each, we need to a follow-up query
			// to pull a full person profile by search pointer
			searchPointer := person.SearchPointer
			var searchResponse *Response
			if searchResponse, err = c.SearchByPointer(ctx, searchPointer); err != nil {
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
func (c *Client) SearchByPointer(ctx context.Context, searchPointer string) (*Response, error) {
	// So we have a search pointer?
	if len(searchPointer) < 20 {
		return nil, ErrInvalidSearchPointer
	}

	// Start the post data
	postData := url.Values{}

	// Add the API key
	postData.Add(fieldAPIKey, c.options.apiKey)

	// Option for pretty response
	if !c.options.searchOptions.Search.Pretty {
		postData.Add(fieldPretty, valueFalse)
	}

	// Add the search pointer
	postData.Add(fieldSearchPointer, searchPointer)

	// Fire the request
	response, err := httpRequest(ctx, c, searchAPIEndpoint, &postData)
	if err != nil {
		return nil, err
	} else if len(response.Error) > 0 {
		return nil, fmt.Errorf("%w: %s", ErrAPIResponse, response.Error)
	}
	return response, nil
}
