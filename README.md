# httpclient

The httpclient module provides a simple and flexible way to interact with HTTP calls. It is designed to be lightweight, easy to use, and customizable to fit your specific needs.

## Installation

To use httpclient, you need to have Go installed and set up on your system. Once you have Go installed, you can install http-client-go using the following command:

```bash
go get -u https://github.com/BusquetsLA/httpclient
```

## Usage
httpclient is designed to be easy to use, with sensible defaults and options that can be easily overridden as needed. Here's an example of how to create an HTTP client and make a request:

```go
package main

import (
	"errors"
	"fmt"

	"github.com/BusquetsLA/httpclient/httpgo"
)

var (
	httpClient = httpgo.New().Build()
)

type Endpoints struct {
	CurrentUserUrl string `json:"current_user_url"`
	UserUrl        string `json:"user_url"`
	RepositoryUrl  string `json:"repository_url"`
}

// Create a new HTTP Get request.
func Get() (*Endpoints, error) {
	// Make the request and handle the response.
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 299 {
		return nil, errors.New("error when trying to fetch github endpoints")
	}
	
	var endpoints Endpoints
	// Unmarshal the response and use it as needed.
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	return &endpoints, nil
}
```

## Contributing

Pull requests are welcome! If you find a bug or have a feature request, please open an issue first
to discuss what you would like to change. If you would like to contribute code, please fork the repository and submit a pull request.

Please make sure to update tests as appropriate.
