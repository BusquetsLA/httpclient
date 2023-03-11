package examples

import (
	"errors"
	"net/http"
)

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private"`
}

type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

func CreateRepository(req Repository) (*Repository, error) {
	res, err := httpClient.Post("https://api.github.com/user/repos", req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		var githubError GithubError
		if err := res.UnmarshalJson(&githubError); err != nil {
			return nil, errors.New("error processing github error response when creating a new repository")
		}
		return nil, errors.New(githubError.Message)
	}

	var response Repository
	if err := res.UnmarshalJson(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
