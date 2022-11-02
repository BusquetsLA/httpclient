package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// httpMethod := "GET"
	url := "https://api.github.com"

	client := http.Client{}

	response, err := client.Get(url)
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
