package httpgo

import (
	"net/http"
	"time"
)

// clientBuilder structure provides a way to configure the client, storing a previous client, the headers and other values
type clientBuilder struct {
	headers      http.Header // default headers from http pkg
	connTimeout  time.Duration
	resTimeout   time.Duration
	maxIdleConns int
	disTimeout   bool
	client       *http.Client
}

type ClientBuilder interface {
	// methods for configuration, should be defined only when creating the client
	Build() Client
	SetHeaders(headers http.Header) ClientBuilder
	SetConnTimeout(timeout time.Duration) ClientBuilder
	SetResTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConns(maxConns int) ClientBuilder
	DisableTimeouts(disTimeouts bool) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
}

func New() ClientBuilder { // single http client being used every time for every request
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

// SetHeaders func configures the headers for the client, besides the default headers
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

// SetConnTimeout func configures the connection timeout for the client
func (c *clientBuilder) SetConnTimeout(timeout time.Duration) ClientBuilder {
	c.connTimeout = timeout
	return c
}

// SetResTimeout func configures the response timeout for the client
func (c *clientBuilder) SetResTimeout(timeout time.Duration) ClientBuilder {
	c.resTimeout = timeout
	return c
}

// SetMaxIdleConns func configures the maximum number of idle connections for the client
func (c *clientBuilder) SetMaxIdleConns(maxConns int) ClientBuilder {
	c.maxIdleConns = maxConns
	return c
}

// DisableTimeouts func configures the disabling of timeouts for the client
func (c *clientBuilder) DisableTimeouts(disTimeout bool) ClientBuilder {
	c.disTimeout = disTimeout
	return c
}

// SetHttpClient func allows to use the configuration of a prevous client in a new one
func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}
