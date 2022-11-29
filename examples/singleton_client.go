package examples

import (
	"github.com/BusquetsLA/httpclient/httpgo"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() httpgo.Client {
	client := httpgo.New(). // here goes all the configurations you would like to add
				Build() // and finish it with build
	return client

	// currentClient := http.Client{}
	// client := httpgo.New().
	// 	SetResTimeout(3 * time.Second).
	// 	SetConnTimeout(3 * time.Second).
	// 	SetHttpClient(&currentClient).
	// 	Build()
	// return client
}
