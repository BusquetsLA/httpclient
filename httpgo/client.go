package httpgo

import (
	"net/http"
	"sync"
)

type httpClient struct {
	client     *http.Client
	builder    *clientBuilder
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

func getHeaders(headers ...http.Header) http.Header {
	// variadic functions can be called with any number of trailing arguments, but the variadic arg always has to come last
	// this checks the headers don't come empty and if so it fills them with default headers
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

// HTTP Call Methods:
func (c *httpClient) Get(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, nil, getHeaders(headers...))
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, body, getHeaders(headers...))
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, body, getHeaders(headers...))
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, body, getHeaders(headers...))
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, nil, getHeaders(headers...))
}

func (c *httpClient) Options(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, nil, getHeaders(headers...))
}
