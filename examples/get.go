package examples

import (
	"fmt"
)

type Endpoints struct {
	CurrentUserUrl string `json:"current_user_url"`
	FollowersUrl   string `json:"followers_url"`
	FollowingUrl   string `json:"following_url"`
	UserUrl        string `json:"user_url"`
	RepositoryUrl  string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) { // Get
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("status code recieved: %v, %v, response body: %v", response.Status, response.StatusCode, response.String())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Printf("repository URL: %s", endpoints.RepositoryUrl)
	return &endpoints, nil
}
