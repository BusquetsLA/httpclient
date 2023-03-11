package core

import "net/http"

// HttpClient is the interface for the http client implementation used by the httpgo package.
// Any implementation of this interface can be used to send http requests.
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
