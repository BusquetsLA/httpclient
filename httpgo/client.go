package httpgo

import (
	"net/http"
	"time"
)

type httpClient struct {
	client       *http.Client
	Headers      http.Header // default headers from http pkg
	connTimeout  time.Duration
	resTimeout   time.Duration
	maxIdleConns int
}

func New() HttpClient { // single http client being used every time for every request
	httpClient := &httpClient{}
	return httpClient
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	SetConnTimeout(timeout time.Duration)
	SetResTimeout(timeout time.Duration)
	SetMaxIdleConns(maxConns int)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) SetConnTimeout(timeout time.Duration) {
	c.connTimeout = timeout
}

func (c *httpClient) SetResTimeout(timeout time.Duration) {
	c.resTimeout = timeout
}

func (c *httpClient) SetMaxIdleConns(maxConns int) {
	c.maxIdleConns = maxConns
}

// HTTP Methods:

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
