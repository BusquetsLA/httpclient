package httpgo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/BusquetsLA/httpclient/core"
	"github.com/BusquetsLA/httpclient/mock"
)

const (
	defaultMaxIdleConn = 5
	defaultResTimeout  = 5 * time.Second
	defaultConnTimeout = 1 * time.Second
)

// do is a private method of httpClient that makes the actual HTTP request
// do performs an HTTP request with the specified HTTP method and URL.
// It sets the request headers and body, sends the request, and returns
// the response or an error if something goes wrong.
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
	// Get the request headers to be used for the HTTP request
	// moved here to have acccess to the headers before creating the request to make the request body
	reqHeaders := c.getRequestHeaders(headers)

	// Get the request body to be used for the HTTP request
	reqBody, err := c.getRequestBody(reqHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
		// return nil, errors.New("unable to create request body")
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
		// return nil, errors.New("unable to create new request")
	}

	// Set the request headers
	req.Header = reqHeaders

	// Make the HTTP request using the underlying HTTP client
	res, err := c.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	// Defering the close of the body until returning it, then closing the response body when done reading it
	defer res.Body.Close()
	// Read the response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Create a response object containing the response status code, headers, and body
	response := core.Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       resBody,
	}

	return &response, nil
}

// getHttpClient returns the HTTP client to be used for making HTTP requests
// getHttpClient returns an HTTP client instance. If a mock server is enabled,
// it returns the mocked client. Otherwise, it creates a new HTTP client with
// the specified timeouts, max idle connections per host, and other settings.
func (c *httpClient) getHttpClient() core.HttpClient {
	// If MockupServer is enabled, use the mocked HTTP client instead, if not the library will make the real call to the api
	if mock.MockupServer.IsEnabled() {
		return mock.MockupServer.GetMockedClient()
	}

	// Create a new HTTP client only once, making the client concurrent safe
	c.clientOnce.Do(func() {
		// If a custom client has already been set, use it instead of creating a new one
		if c.builder.client != nil {
			c.client = c.builder.client
			return // if there is a client already built it will miss the c.client = &http.Client{} to build one
		}

		// Otherwise, build an HTTP client with the specified configuration for timeout and transport settings.
		c.client = &http.Client{
			Timeout: c.getConnTimeout() + c.getResTimeout(), // To configure the overall client timeout
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConn(), // This number should be based solely on the traffic pattern that you have in your application
				ResponseHeaderTimeout: c.getResTimeout(),  // Max amount of time to wait for a response when a request is sent
				DialContext: (&net.Dialer{
					Timeout: c.getConnTimeout(),
				}).DialContext, // To set max amount of time to wait for a given connection
			},
		}
	})

	return c.client
}

// BODY

// getRequestBody marshals the request body into JSON or XML format, depending on
// the specified content type, and returns the request body to be used for the HTTP request.
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

// HEADERS

// getRequestHeaders returns the HTTP headers for the request. It sets the
// standard headers, custom headers, and user agent header, if available.
func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// Adding standard headers for every method.
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Adding custom headers for every method.
	for header, value := range headers {
		if len(value) > 0 { // to avoid a panic if header comes empty
			result.Set(header, value[0])
		}
	}

	// Setting User-Agent header if it's not already configured.
	if c.builder.userAgent != "" {
		if result.Get("User-Agent") != "" {
			result.Set("User-Agent", c.builder.userAgent)
		}
	}

	return result
}

func getHeaders(headers ...http.Header) http.Header {
	// variadic functions can be called with any number of trailing arguments, but the variadic arg always has to come last
	// this checks the headers don't come empty and if so it fills them with default headers
	if len(headers) > 0 {
		return headers[0]
	}

	return http.Header{}
}

// TIMEOUTS & IDLE CONNECTIONS

// getMaxIdleConn returns the maximum number of idle connections that the HTTP client will keep alive.
// If the builder specifies a maximum number of idle connections, that value is returned.
// Otherwise, the defaultMaxIdleConn constant is returned.
func (c *httpClient) getMaxIdleConn() int {
	// check if a maximum number of idle connections is specified
	if c.builder.maxIdleConns > 0 {
		return c.builder.maxIdleConns
	}
	// return the default maximum number of idle connections if no value is specified
	return defaultMaxIdleConn
}

// getResTimeout returns the maximum amount of time that the HTTP client will wait for a response from the server.
// If the builder specifies a response timeout, that value is returned.
// If the builder has disabled timeouts, this function returns 0.
// Otherwise, the defaultResTimeout constant is returned.
func (c *httpClient) getResTimeout() time.Duration {
	// check if timeouts are disabled
	if c.builder.disTimeout {
		return 0
	}
	// check if a response timeout is specified
	if c.builder.resTimeout > 0 {
		return c.builder.resTimeout
	}
	// return the default timeout if no value is specified
	return defaultResTimeout
}

// getConnTimeout returns the maximum amount of time that the HTTP client will wait for a connection to be established.
// If the builder specifies a connection timeout, that value is returned.
// If the builder has disabled timeouts, this function returns 0.
// Otherwise, the defaultConnTimeout constant is returned.
func (c *httpClient) getConnTimeout() time.Duration {
	// check if timeouts are disabled
	if c.builder.disTimeout {
		return 0
	}
	// check if a connection timeout is specified
	if c.builder.connTimeout > 0 {
		return c.builder.connTimeout
	}
	// return the default timeout if no value is specified
	return defaultConnTimeout
}
