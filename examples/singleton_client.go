package examples

import (
	"net/http"
	"time"

	"github.com/BusquetsLA/httpclient/httpgo"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() httpgo.Client {
	// client := httpgo.New(). // here goes all the configurations you would like to add
	// 			Build() // and finish it with build
	// return client

	headers := make(http.Header)

	// currentClient := http.Client{}
	client := httpgo.New().
		// SetHttpClient(&currentClient).
		SetHeaders(headers).
		SetResTimeout(3 * time.Second).
		SetConnTimeout(3 * time.Second).
		Build()
	return client
}
