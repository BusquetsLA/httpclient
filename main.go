package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpMethod := "GET"
	url := "https://api.github.com"

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, nil)

	request.Header.Set("Accept", "application/json") // The Accept header tells the server the response content type you can understand.

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // A defer statement defers the execution of a function until the surrounding function returns.
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status code recieved: %v, Response Body: %v", response.StatusCode, string(bytes))
}
