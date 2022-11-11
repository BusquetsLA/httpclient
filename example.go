package main

import (
	"fmt"
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
	// lauti := User{"lauti", "busquets"}

	for i := 0; i < 10; i++ {
		go func() {
			getGithubUrls()
		}()
	}
}

func getGithubUrls() {
	response, err := githubClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	lauti := User{"lauti", "busquets"}
	if err := response.JsonUnmarshal(&lauti); err != nil {
		panic(err)
	}
	fmt.Println(lauti.FirstName)

	// bytes, err := ioutil.ReadAll(response.Body) // no more need for this
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("status code recieved: %v, response body: %v", response.StatusCode(), response.String())
}

func getGithubClient() httpgo.Client {
	client := httpgo.New().
		SetConnTimeout(2 * time.Second).
		DisableTimeouts(true).
		Build() // to create a client with all the configuration from the begining, has to end with Build()

	return client
}
