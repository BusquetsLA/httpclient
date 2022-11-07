package httpgo

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) { // rule of thumb for % of coverage: 1 test case for every return that the function has
	// Initialization:
	client := httpClient{}
	requestHeaders := make(http.Header)
	requestHeaders.Set("Content-Type", "application/json") // 1
	requestHeaders.Set("User-Agent", "BusquetsLA")         // 2
	client.Headers = requestHeaders

	// Execution:
	headers := make(http.Header)
	headers.Set("Accept", "*/*")                      // 3
	finalHeaders := client.getRequestHeaders(headers) // the final list with every added header

	// Validation:
	want := 3 // amount of headers set in the test
	if got := len(finalHeaders); got != want {
		t.Error("expected 3 headers") // this can still be HIGHLY improved
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type recieved")
	}
	if finalHeaders.Get("User-Agent") != "BusquetsLA" {
		t.Error("invalid user agent recieved")
	}
	if finalHeaders.Get("Accept") != "*/*" {
		t.Error("invalid accept recieved")
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
