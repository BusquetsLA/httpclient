package httpgo

import (
	"net/http"
	"sync"

	"github.com/BusquetsLA/httpclient/core"
)

type httpClient struct {
	client     *http.Client
	builder    *clientBuilder
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
	Options(url string, headers ...http.Header) (*core.Response, error)
}

// HTTP Call Methods:
func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, nil, getHeaders(headers...))
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, body, getHeaders(headers...))
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, body, getHeaders(headers...))
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, body, getHeaders(headers...))
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, nil, getHeaders(headers...))
}

func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, nil, getHeaders(headers...))
}
