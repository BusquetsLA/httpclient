package examples

import (
	"errors"
	"net/http"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

func CreateRepository(req Repository) (*Repository, error) {
	// bytes, _ := json.Marshal(req)
	// fmt.Println(string(bytes))
	// fmt.Println(string(bytes))

	res, err := httpClient.Post("https://api.github.com/user/repos", req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusCreated {
		var gitubError GithubError
		if err := res.JsonUnmarshal(&gitubError); err != nil {
			return nil, errors.New("error processing github error response when creating new repository")
		}
		return nil, errors.New(gitubError.Message)
	}

	var newRepository Repository
	if err := res.JsonUnmarshal(&newRepository); err != nil {
		return nil, err
	}

	return &newRepository, nil
}
