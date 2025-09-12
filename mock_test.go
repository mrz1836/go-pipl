package pipl

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// validResponse will return valid response(s)
type validResponse struct{}

// Do will do the HTTP request
func (v *validResponse) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, ErrMissingRequest
	}

	// Parse the form data
	if err := req.ParseForm(); err != nil {
		return nil, err
	}

	// log.Println("encode: ", req.Form.Encode())

	formEncoded := req.Form.Encode()

	// Search via Username (all options)
	if formEncoded == "key="+testKey+"&person=%7B%22usernames%22%3A%5B%7B%22content%22%3A%22superman%40facebook%22%7D%5D%7D&pretty=false&show_sources=all" ||
		formEncoded == "hide_sponsored=true&infer_persons=true&key="+testKey+"&live_feeds=false&match_requirements=email&minimum_match=0.1&person=%7B%22usernames%22%3A%5B%7B%22content%22%3A%22superman%40facebook%22%7D%5D%7D&show_sources=all&source_category_requirements=professional_and_business&top_match=true" ||
		formEncoded == "key="+testKey+"&pretty=false&search_pointer="+testSearchPointer ||
		formEncoded == "key="+testKey+"&search_pointer="+testSearchPointer {
		response, err := loadResponseData("response_success.json")
		if err != nil {
			return nil, err
		}
		var b []byte
		if b, err = json.Marshal(response); err != nil {
			return nil, err
		}
		resp.StatusCode = http.StatusOK
		resp.Body = io.NopCloser(bytes.NewBuffer(b))
		return resp, nil
	}

	// No request found, return an error
	resp.Body = io.NopCloser(bytes.NewReader([]byte(`{"error":"no-route-found"}`)))
	return resp, ErrRequestNotFound
}

// errorHTTPResponse will return error response(s)
type errorHTTPResponse struct{}

// Do will do the HTTP request
func (v *errorHTTPResponse) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest
	return resp, ErrBadRequest
}

// errorBadJSONResponse will return error response(s)
type errorBadJSONResponse struct{}

// Do will do the HTTP request
func (v *errorBadJSONResponse) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusOK
	resp.Body = io.NopCloser(bytes.NewReader([]byte(`{error:bad-json}`)))
	return resp, nil
}

// errorMissingAPIKeyResponse will return error response(s)
type errorMissingAPIKeyResponse struct{}

// Do will do the HTTP request
func (v *errorMissingAPIKeyResponse) Do(_ *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusForbidden
	resp.Body = io.NopCloser(bytes.NewReader([]byte(`{"@http_status_code": 403,"error": "Please provide an API key"}`)))
	return resp, nil
}
