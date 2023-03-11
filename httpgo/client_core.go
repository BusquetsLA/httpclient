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

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
	reqHeaders := c.getRequestHeaders(headers) // moved here to have acccess to the headers before creating the request to make the request body

	reqBody, err := c.getRequestBody(reqHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
		// return nil, errors.New("unable to create request body")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
		// return nil, errors.New("unable to create new request")
	}

	req.Header = reqHeaders

	res, err := c.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close() // defering the close of the body until returning it
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := core.Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       resBody,
	}

	return &response, nil
}

func (c *httpClient) getHttpClient() core.HttpClient {
	if mock.MockupServer.IsEnabled() { // if there isn't a mock the library will make the real call to the api
		return mock.MockupServer.GetMockedClient()
	}

	c.clientOnce.Do(func() { // to make the client concurrent safe
		if c.builder.client != nil {
			c.client = c.builder.client
			return // if there is a client already built it will miss the c.client = &http.Client{} to build one
		}
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

// BODY
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

	// set User-Agent header if it's not configured aready
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
func (c *httpClient) getMaxIdleConn() int {
	if c.builder.maxIdleConns > 0 {
		return c.builder.maxIdleConns
	}

	return defaultMaxIdleConn
}

func (c *httpClient) getResTimeout() time.Duration {
	if c.builder.disTimeout { // now it checks for disable timeouts first
		return 0
	}

	if c.builder.resTimeout > 0 {
		return c.builder.resTimeout
	}

	return defaultResTimeout
}

func (c *httpClient) getConnTimeout() time.Duration {
	if c.builder.disTimeout { // now it checks for disable timeouts first
		return 0
	}

	if c.builder.connTimeout > 0 {
		return c.builder.connTimeout
	}

	return defaultConnTimeout
}
