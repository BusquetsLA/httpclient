package examples

import (
	"github.com/BusquetsLA/httpclient/httpgo"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() httpgo.Client {
	client := httpgo.New().
		Build()
	return client
}
