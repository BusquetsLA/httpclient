package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BusquetsLA/httpclient/httpgo"
)

var (
	githubClient = getGithubClient()
)

type User struct {
	FirstName string `json:"first_name"` //fields in lowercase will never be exported from the "example" package
	LastName  string `json:"last_name"`
}

func main() {
	lauti := User{"lauti", "busquets"}
	createUser(lauti)
	getGithubUrls()
}

func createUser(user User) {
	response, err := githubClient.Post("https://api.github.com", nil, user) // ofc returns 404
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status code recieved: %v, response body: %v", response.StatusCode, string(bytes))
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

func getGithubClient() httpgo.HttpClient {
	client := httpgo.New()

	// client.SetConnTimeout(2 * time.Second)
	// client.SetResTimeout(50 * time.Millisecond) // this would return timeout

	standardHeaders := make(http.Header)
	standardHeaders.Set("Accept", "application/json")
	client.SetHeaders(standardHeaders)

	return client
}
