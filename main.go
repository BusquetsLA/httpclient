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

	fmt.Println(response.StatusCode)

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
