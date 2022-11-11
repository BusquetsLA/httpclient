package httpgo

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status     string
	statusCode int
	headers    http.Header
	body       []byte
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Headers() http.Header {
	return r.headers
}

// different types of body return:
func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) JsonUnmarshal(data interface{}) error {
	return json.Unmarshal(r.Bytes(), data)
}
