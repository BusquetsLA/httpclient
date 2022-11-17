package httpgo

import (
	"net/http"
	"testing"
	"time"
)

func TestGetRequestHeaders(t *testing.T) { // rule of thumb for % of coverage: 1 test case for every return that the function has
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
	client.builder = &clientBuilder{
		headers: requestHeaders,
	}

	finalHeaders := client.getRequestHeaders(requestHeaders) // the final list with every added header

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

func TestGetResTimeout(t *testing.T) {
	client := httpClient{}
	client.builder = &clientBuilder{}
	t.Run("DefaultResponseTimeout", func(t *testing.T) {
		timeout := client.getResTimeout()
		if timeout != defaultResTimeout {
			t.Error("expected default timeout")
		}
	})
	t.Run("CustomResponseTimeout", func(t *testing.T) {
		client.builder = &clientBuilder{
			resTimeout: 3 * time.Second,
		}
		timeout := client.getResTimeout()
		if timeout != 3*time.Second {
			t.Errorf("expected %v, got %v", 3*time.Second, timeout)
		}
	})
	t.Run("DisabledResponseTimeout", func(t *testing.T) {
		client.builder = &clientBuilder{
			disTimeout: true,
		}
		if client.builder.resTimeout != 0 {
			t.Errorf("expected no timeout, got %v", client.builder.resTimeout)
		}
	})
}

func TestGetConnTimeout(t *testing.T) {
	client := httpClient{}
	client.builder = &clientBuilder{}
	t.Run("DefaultConnectionTimeout", func(t *testing.T) {
		timeout := client.getConnTimeout()
		if timeout != defaultConnTimeout {
			t.Error("expected default timeout")
		}
	})
	t.Run("CustomResponseTimeout", func(t *testing.T) {
		client.builder = &clientBuilder{
			connTimeout: 3 * time.Second,
		}
		timeout := client.getConnTimeout()
		if timeout != 3*time.Second {
			t.Errorf("expected %v, got %v", 3*time.Second, timeout)
		}
	})
	t.Run("DisabledResponseTimeout", func(t *testing.T) {
		client.builder = &clientBuilder{
			disTimeout: true,
		}
		if client.builder.resTimeout != 0 {
			t.Errorf("expected no timeout, got %v", client.builder.connTimeout)
		}
	})
}

func TestGetMaxIdleConn(t *testing.T) {
	client := httpClient{}
	client.builder = &clientBuilder{}
	idleConn := client.getMaxIdleConn()
	t.Run("DefaultMaxIdleConnections", func(t *testing.T) {
		if idleConn != defaultMaxIdleConn {
			t.Error("expected default max idle connections")
		}
	})
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

// func TestDo(t *testing.T)             {}
// func TestGetHttpClient(t *testing.T)  {}
