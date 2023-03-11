package httpgo

import (
	"net/http"
	"testing"
	"time"
)

// rule of thumb for % of coverage: 1 test case for every return that the function has

// TestGetRequestHeaders tests the getRequestHeaders method of httpClient.

func TestGetRequestHeaders(t *testing.T) {
	// Create an instance of httpClient.
	client := httpClient{}

	// Define the tests with a slice of structs containing the key and value for each header.
	tests := []struct {
		name string // The key of the header.
		want string // The value of the header.
	}{
		{"Content-Type", "application/json"},
		{"User-Agent", "BusquetsLA"},
		{"Accept", "*/*"},
	}

	// Create an HTTP header to store the headers for the request.
	requestHeaders := make(http.Header)
	// Set each header in the HTTP header.
	for _, tt := range tests {
		requestHeaders.Set(tt.name, tt.want)
	}

	// Set the request headers for the client.
	client.builder = &clientBuilder{
		headers: requestHeaders,
	}

	// Get the final list of headers with getRequestHeaders.
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Ensure that the finalHeaders list contains all the headers we set in the test.
	want := len(tests)
	if got := len(finalHeaders); got != want {
		t.Errorf("expected %d headers, got %d", want, got)
	}

	// Check that each header is correct in the final list of headers.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalHeaders.Get(tt.name); got != tt.want {
				t.Errorf("invalid %s header recieved", tt.name)
			}
		})
	}
}

// TestGetRequestBody tests the getRequestBody method of httpClient.

func TestGetRequestBody(t *testing.T) {
	// Create an instance of httpClient.
	client := httpClient{}

	// Test for nil body.
	t.Run("nil Body response", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("expected no error when passing nil body")
		}
		if body != nil {
			t.Error("expected nil body when passing nil body")
		}
	})

	// Define the tests with a slice of structs containing the content type and expected request body.
	requestBody := []string{"pauli", "brujita"}
	tests := []struct {
		contentType string // The content type header key.
		want        string // The expected request body for that content type.
	}{
		{"application/json", `["pauli","brujita"]`},
		{"application/xml", `<string>pauli</string><string>brujita</string>`},
		{"default/contentType", `["pauli","brujita"]`},
		{"", `["pauli","brujita"]`},
	}

	// Test each content type with the requestBody.
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

func TestGetResTimeout(t *testing.T) {
	client := httpClient{}
	client.builder = &clientBuilder{}

	// Define the tests with a slice of structs containing the name of each test, the response timeout and disable timeouts and finally, the expected return value.
	tests := []struct {
		name            string
		responseTimeout time.Duration
		disableTimeouts bool
		want            time.Duration
	}{
		// Test case 1: Default Response Timeout scenario
		{"Default Response Timeout: ", client.getResTimeout(), false, defaultResTimeout},
		// Test case 2: Custom Response Timeout scenario
		{"Custom Response Timeout: ", 3 * time.Second, false, 3 * time.Second},
		// Test case 3: Disabled Response Timeout scenario
		{"Disabled Response Timeout: ", client.getResTimeout(), true, 0},
	}

	// Loop through each test scenario and run the test case.
	for _, tt := range tests {
		// Run each test case as a subtest with a unique name and function.
		t.Run(tt.name, func(t *testing.T) {
			// Set the httpClient's builder to a new clientBuilder instance with the current test case input values.
			client.builder = &clientBuilder{
				resTimeout: tt.responseTimeout,
				disTimeout: tt.disableTimeouts,
			}
			// Get the response timeout value using the current test case input values.
			timeout := client.getResTimeout()
			// Check if the returned value matches the expected value.
			if timeout != tt.want {
				// If the returned value does not match the expected value, log an error with details.
				t.Errorf("Expected response Timeout of %v, got %v", tt.want, timeout)
			}
		})
	}
}

// TestGetConnTimeout tests the GetConnTimeout method of the httpClient struct.
func TestGetConnTimeout(t *testing.T) {
	client := httpClient{}
	client.builder = &clientBuilder{}

	// Define test cases
	tests := []struct {
		name              string
		connectionTimeout time.Duration
		disableTimeouts   bool
		want              time.Duration
	}{
		{"Default Connection Timeout: ", client.getConnTimeout(), false, defaultConnTimeout},
		{"Custom Connection Timeout: ", 10 * time.Second, false, 10 * time.Second},
		{"Disabled Connection Timeout: ", client.getConnTimeout(), true, 0},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the connection timeout and disableTimeouts properties of client.builder
			client.builder = &clientBuilder{
				connTimeout: tt.connectionTimeout,
				disTimeout:  tt.disableTimeouts,
			}
			// Verify that GetConnTimeout returns the expected timeout value
			if timeout := client.getConnTimeout(); timeout != tt.want {
				t.Errorf("Expected connection Timeout of %v, got %v", tt.want, timeout)
			}
		})
	}
}

// TestGetMaxIdleConn tests the GetMaxIdleConn method of the httpClient struct.
func TestGetMaxIdleConn(t *testing.T) {
	client := httpClient{}
	client.builder = &clientBuilder{}
	idleConn := client.getMaxIdleConn()

	// Verify the default value of max idle connections
	t.Run("DefaultMaxIdleConnections", func(t *testing.T) {
		if idleConn != defaultMaxIdleConn {
			t.Error("expected default max idle connections")
		}
	})

	// Verify a custom value of max idle connections
	t.Run("CustomMaxIdleConnections", func(t *testing.T) {
		client.builder = &clientBuilder{
			maxIdleConns: 10,
		}
		newIdleConn := client.getMaxIdleConn()
		if newIdleConn != 10 {
			t.Errorf("expected %v max idle connections, got %v", 10, idleConn)
		}
	})
}
