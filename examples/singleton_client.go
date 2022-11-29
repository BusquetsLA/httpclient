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
}
