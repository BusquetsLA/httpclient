package httpgo

import (
	"net/http"
	"time"
)

// clientBuilder provides a way to configure the client, storing a previous client, the headers, and other values
type clientBuilder struct {
	headers      http.Header   // HTTP headers
	connTimeout  time.Duration // connection timeout
	resTimeout   time.Duration // response timeout
	maxIdleConns int           // maximum number of idle connections
	disTimeout   bool          // flag indicating if timeouts are disabled
	client       *http.Client  // HTTP client
	userAgent    string        // user agent
}

// ClientBuilder interface defines the methods for configuring the HTTP client, should be defined only when creating the client
type ClientBuilder interface {
	Build() Client
	SetHeaders(headers http.Header) ClientBuilder
	SetConnTimeout(timeout time.Duration) ClientBuilder
	SetResTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConns(maxConns int) ClientBuilder
	DisableTimeouts(disTimeouts bool) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
}

// New returns a new instance of ClientBuilder, a single http client being used every time for every request
func New() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

// Build returns a new instance of Client
func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

// SetHeaders configures the HTTP headers for the client, besides the default headers
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

// SetConnTimeout configures the connection timeout for the client
func (c *clientBuilder) SetConnTimeout(timeout time.Duration) ClientBuilder {
	c.connTimeout = timeout
	return c
}

// SetResTimeout configures the response timeout for the client
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

// SetUserAgent func configures the user agent for the client
func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}
