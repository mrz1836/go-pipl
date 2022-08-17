package pipl

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// httpRequest is a generic pipl request wrapper that can be used without the constraints
// of the Search or SearchByPointer methods
func httpRequest(ctx context.Context, client *Client, endpoint string,
	params *url.Values) (response *Response, err error) {

	// Start the request
	var request *http.Request
	if request, err = http.NewRequestWithContext(
		ctx, http.MethodPost, endpoint, strings.NewReader(params.Encode()),
	); err != nil {
		return
	}

	// Set the headers
	request.Header.Set("User-Agent", client.options.userAgent)

	// Set the content type on method
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Fire the http request
	var resp *http.Response
	if resp, err = client.options.httpClient.Do(request); err != nil {
		return
	}

	// Close the response body
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	// Read the body
	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	// Parse the response
	response = new(Response)
	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// Thumbnail generation enabled?
	if client.options.searchOptions.Thumbnail.Enabled {

		// Process the current person
		response.Person.ProcessThumbnails(client.options.searchOptions.Thumbnail)

		// Do we have possible persons?
		if len(response.PossiblePersons) > 0 {
			for index := range response.PossiblePersons {
				response.PossiblePersons[index].ProcessThumbnails(client.options.searchOptions.Thumbnail)
			}
		}
	}

	return
}
