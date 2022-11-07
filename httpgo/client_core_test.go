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

	// Execution
	headers := make(http.Header)
	headers.Set("Accept", "*/*") // 3

	finalHeaders := client.getRequestHeaders(headers) // the final list with every added header

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("expected 3 headers") // this can be HIGHLY improved
	}
}
