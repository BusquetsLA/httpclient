package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// different types of body return:
func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(data interface{}) error {
	return json.Unmarshal(r.Bytes(), data)
}
