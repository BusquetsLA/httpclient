package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BusquetsLA/httpclient/httpgo"
)

var (
	// githubClient = httpgo.New()
	githubClient = getGithubClient()
)

func main() {
	getGithubUrls()
}

func getGithubUrls() {
	response, err := githubClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status code recieved: %v, response body: %v", response.StatusCode, string(bytes))
}

func getGithubClient() httpgo.HttpClient { // just to set the headers
	client := httpgo.New()
	standardHeaders := make(http.Header)
	standardHeaders.Set("Accept", "application/json")
	client.SetHeaders(standardHeaders)
	return client
}
