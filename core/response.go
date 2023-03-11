package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string      // HTTP response status
	StatusCode int         // HTTP response status code
	Headers    http.Header // HTTP response headers
	Body       []byte      // HTTP response body
}

// Bytes returns the HTTP response body as a byte array
func (r *Response) Bytes() []byte {
	return r.Body
}

// String returns the HTTP response body as a string
func (r *Response) String() string {
	return string(r.Body)
}

// UnmarshalJson unmarshals the JSON-encoded HTTP response body into a Go data structure
func (r *Response) UnmarshalJson(data interface{}) error {
	return json.Unmarshal(r.Bytes(), data)
}
