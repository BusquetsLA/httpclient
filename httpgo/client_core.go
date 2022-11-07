package httpgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	requestHeaders := c.getRequestHeaders(headers) // moved here to have acccess to the headers before creating the request to make the request body

	requestBody, err := c.getRequestBody(requestHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, errors.New("unable to create request body")
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	request.Header = requestHeaders

	return c.client.Do(request)
}

func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// addign standard headers for every method
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// addign custom headers for every method
	for header, value := range headers {
		if len(value) > 0 { // to avoid a panic if header comes empty
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch contentType {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body) // TODO: add more cases
	}
}
