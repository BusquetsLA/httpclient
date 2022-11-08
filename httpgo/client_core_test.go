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
	client := httpClient{}

	// first for nil body
	t.Run("nil Body response", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("expected no error when passing nil body")
		}
		if body != nil {
			t.Error("expected nil body when passing nil body")
		}
	})

	// then the rest of test cases
	requestBody := []string{"pauli", "brujita"}
	tests := []struct {
		contentType string // key of the header
		want        string // value of the header
	}{
		// {"", nil}
		{"application/json", `["pauli","brujita"]`},
		{"application/xml", `<string>pauli</string><string>brujita</string>`},
		{"default/contentType", `["pauli","brujita"]`},
		{"", `["pauli","brujita"]`},
	}
	for _, tt := range tests {
		t.Run(tt.contentType, func(t *testing.T) {
			got, err := client.getRequestBody(tt.contentType, requestBody)
			// expected := (string(got)) // to see what we should be getting
			if err != nil {
				t.Errorf("expected no error when passing %s body for %s content type", tt.want, tt.contentType)
			}
			if string(got) != (string(tt.want)) {
				t.Errorf("expected %s body when passing that content type", tt.contentType)
			}
		})
	}
}
