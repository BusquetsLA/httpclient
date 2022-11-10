package httpgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"time"
)

const (
	defaultMaxIdleConn = 5
	defaultResTimeout  = 5 * time.Second
	defaultConnTimeout = 1 * time.Second
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	requestHeaders := c.getRequestHeaders(headers) // moved here to have acccess to the headers before creating the request to make the request body

	requestBody, err := c.getRequestBody(requestHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, errors.New("unable to create request body")
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	request.Header = requestHeaders
	client := c.getHttpClient()

	return client.Do(request)
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() { // to make the client concurrent safe
		c.client = &http.Client{
			Timeout: c.getConnTimeout() + c.getResTimeout(), // to configure the overall client timeout
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConn(), // this number should be based solely on the traffic pattern that you have in your application
				ResponseHeaderTimeout: c.getResTimeout(),  // max amount of time to wait for a response when a request is sent
				DialContext: (&net.Dialer{
					Timeout: c.getConnTimeout(),
				}).DialContext, // to set max amount of time to wait for a given connection
			},
		}
	})

	return c.client
}

func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// addign standard headers for every method
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// addign custom headers for every method
	for header, value := range headers {
		if len(value) > 0 { // to avoid a panic if header comes empty
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch contentType {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body) // TODO: add more cases
	}
}

func (c *httpClient) getMaxIdleConn() int {
	if c.builder.maxIdleConns > 0 {
		return c.builder.maxIdleConns
	}
	return defaultMaxIdleConn
}

func (c *httpClient) getResTimeout() time.Duration {
	if c.builder.resTimeout > 0 {
		return c.builder.resTimeout
	}
	if c.builder.disTimeout {
		return 0
	}
	return defaultResTimeout
}

func (c *httpClient) getConnTimeout() time.Duration {
	if c.builder.connTimeout > 0 {
		return c.builder.connTimeout
	}
	if c.builder.disTimeout {
		return 0
	}
	return defaultConnTimeout
}
