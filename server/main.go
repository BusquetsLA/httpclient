package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(http.ResponseWriter, *http.Request) {
	fmt.Printf("I AM THE SERVER ") // go run server/main.go && curl localhost:8080
}
