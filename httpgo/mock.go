package httpgo

import (
	"fmt"
	"net/http"
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
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d, %s", m.ResStatusCode, http.StatusText(m.ResStatusCode)), // interesante
		statusCode: m.ResStatusCode,
		body:       []byte(m.ResBody),
	}

	return &response, nil
}
