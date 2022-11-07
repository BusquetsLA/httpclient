package httpgo

import (
	"net"
	"net/http"
	"time"
)

type httpClient struct {
	client  *http.Client
	Headers http.Header // default headers from http pkg
}

func New() HttpClient { // single http client being used every time for every request
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,                                                   // this number should be based solely on the traffic pattern that you have in your application
			ResponseHeaderTimeout: 5 * time.Second,                                     // max amount of time to wait for a response when a request is sent
			DialContext:           (&net.Dialer{Timeout: 1 * time.Second}).DialContext, // to set max amount of time to wait for a given connection
		},
	}

	httpClient := &httpClient{
		client: &client,
	}

	return httpClient
}

type HttpClient interface {
	SetHeaders(headers http.Header)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

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
