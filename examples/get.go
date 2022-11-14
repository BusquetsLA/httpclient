package examples

import (
	"fmt"
)

/*
    "current_user_url": "https://api.github.com/user",
	"followers_url": "https://api.github.com/user/followers",
	"following_url": "https://api.github.com/user/following{/target}",
	"user_url": "https://api.github.com/users/{user}",
    "repository_url": "https://api.github.com/repos/{owner}/{repo}",
*/
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
		return nil, err // deal with errors as needed
	}
	fmt.Printf("status code recieved: %v, %v, response body: %v", response.Status(), response.StatusCode(), response.String())

	var endpoints Endpoints
	if err := response.JsonUnmarshal(&endpoints); err != nil {
		return nil, err
	}

	fmt.Printf("Repository URL: %s", endpoints.RepositoryUrl)
	return &endpoints, nil
}
