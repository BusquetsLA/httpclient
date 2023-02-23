package mock

import (
	"fmt"
	"net/http"

	"github.com/BusquetsLA/httpclient/core"
)

// Mock structure provides a way to configure HTTP methods based on the combination between request method, URL and request body
type Mock struct {
	Method        string
	Url           string
	ReqBody       string
	ResBody       string
	Headers       string
	Error         error
	ResStatusCode int
}

// GetResponse returns a Response object based on the mock configuration
func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := core.Response{
		Status:     fmt.Sprintf("%d, %s", m.ResStatusCode, http.StatusText(m.ResStatusCode)), // interesante
		StatusCode: m.ResStatusCode,
		Body:       []byte(m.ResBody),
	}

	return &response, nil
}
