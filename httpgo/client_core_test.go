package httpgo

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) { // rule of thumb for % of coverage: 1 test case for every return that the function has
	// Initialization:
	client := httpClient{}

	tests := []struct {
		name string // key of the header
		want string // value of the header
	}{
		{"Content-Type", "application/json"},
		{"User-Agent", "BusquetsLA"},
		{"Accept", "*/*"},
	}
	requestHeaders := make(http.Header)
	for _, tt := range tests {
		requestHeaders.Set(tt.name, tt.want)
	}

	// Execution:
	finalHeaders := client.getRequestHeaders(requestHeaders) // the final list with every added header

	// Validation:
	want := len(tests) // amount of headers set in the test
	if got := len(finalHeaders); got != want {
		t.Errorf("expected %d headers, got %d", want, got)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalHeaders.Get(tt.name); got != tt.want {
				t.Errorf("invalid %s header recieved", tt.name)
			}
		})
	}
}

func TestGetRequestBody(t *testing.T) {
	// Initialization:
	client := httpClient{}

	t.Run("nil Body response: ", func(t *testing.T) {
		// Execution:
		body, err := client.getRequestBody("", nil)
		// Validation:
		if err != nil {
			t.Error("expected no error when passing nil body")
		}
		if body != nil {
			t.Error("expected nil body when passing nil body")
		}
	})
	t.Run("JSON Body response: ", func(t *testing.T) {
		requestBody := []string{"pauli", "brujita"}
		body, err := client.getRequestBody("application/json", requestBody)

		expected := (string(body))

		if err != nil {
			t.Error("expected no error when marshaling slice as JSON")
		}
		if string(body) != `["pauli","brujita"]` {
			t.Errorf("expected %v, invalid JSON body obtained", expected)
		}
	})
	t.Run("XML Body response: ", func(t *testing.T) {
		requestBody := []string{"pauli", "brujita"}
		body, err := client.getRequestBody("application/xml", requestBody)
		expected := (string(body))

		if err != nil {
			t.Error("expected no error when marshaling slice as JSON")
		}
		if string(body) != `<string>pauli</string><string>brujita</string>` {
			t.Errorf("expected %v, invalid XML body obtained", expected)
		}
	})
	t.Run("default Body response: ", func(t *testing.T) {
		requestBody := []string{"pauli", "brujita"}
		body, err := client.getRequestBody("content/type", requestBody)
		expected := (string(body))

		if err != nil {
			t.Error("expected no error when marshaling slice as JSON")
		}
		if string(body) != `["pauli","brujita"]` {
			t.Errorf("expected %v, invalid JSON body obtained", expected)
		}
	})
}
