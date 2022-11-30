package examples

import (
	"fmt"
	"testing"
)

type Repository struct {
	Name string `json:"name"`
}

func TestPost(t *testing.T) {
	repo := Repository{
		Name: "testing-repo",
	}
	response, err := httpClient.Post("https://api.github.com", repo) // headers now become optional

	fmt.Println(err)
	fmt.Println(response)
}
