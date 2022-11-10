package main

import (
	"fmt"
	"io/ioutil"
	"time"

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

	for i := 0; i < 10; i++ {
		go func() {
			getGithubUrls()
		}()
	}
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

func getGithubClient() httpgo.Client {
	client := httpgo.New().
		SetConnTimeout(2 * time.Second).
		DisableTimeouts(true).
		Build() // to create a client with all the configuration from the begining, has to end with Build()

	// builder := httpgo.New()

	// builder.DisableTimeouts(true)
	// builder.SetConnTimeout(2 * time.Second)
	// builder.SetResTimeout(50 * time.Millisecond) // this would return timeout
	// // builder.SetHeaders(standardHeaders)

	// standardHeaders := make(http.Header)
	// standardHeaders.Set("Accept", "application/json")

	return client
}
