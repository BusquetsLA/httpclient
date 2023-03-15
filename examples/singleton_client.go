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

	headers := make(http.Header)

	// currentClient := http.Client{}
	client := httpgo.New(). // After this goes all the configurations you would like to add to your client
		// SetHttpClient(&currentClient).
		SetHeaders(headers).
		SetResTimeout(3 * time.Second).
		SetConnTimeout(3 * time.Second).
		Build() // And finish it with build
	return client
}
