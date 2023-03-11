package httpgo

import (
	"net/http"
	"sync"

	"github.com/BusquetsLA/httpclient/core"
)

// httpClient is a struct that holds a pointer to an http.Client and a clientBuilder.
// clientOnce is used to ensure that the http.Client is only created once.
type httpClient struct {
	client     *http.Client
	builder    *clientBuilder
	clientOnce sync.Once
}

// Client is an interface that specifies the methods for making HTTP requests.
type Client interface {
	Get(url string, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
	Options(url string, headers ...http.Header) (*core.Response, error)
}

// Get sends a GET request to the specified URL and returns a Response object.
// headers is an optional parameter for specifying additional headers to include in the request.
func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

// Post sends a POST request to the specified URL with the provided body and returns a Response object.
// headers is an optional parameter for specifying additional headers to include in the request.
func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

// Put sends a PUT request to the specified URL with the provided body and returns a Response object.
// headers is an optional parameter for specifying additional headers to include in the request.
func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

// Patch sends a PATCH request to the specified URL with the provided body and returns a Response object.
// headers is an optional parameter for specifying additional headers to include in the request.
func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}

// Delete sends a DELETE request to the specified URL and returns a Response object.
// headers is an optional parameter for specifying additional headers to include in the request.
func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

// Options sends an OPTIONS request to the specified URL and returns a Response object.
// headers is an optional parameter for specifying additional headers to include in the request.
func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
