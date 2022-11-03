package main

import (
	"fmt"
	"net/http"

	"github.com/BusquetsLA/httpclient/httpgo"
)

func main() {
	client := httpgo.New()

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer BusquetsLA")

	reponse, err := client.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response status code: %v", reponse.StatusCode)
}
