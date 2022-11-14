package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/BusquetsLA/httpclient/httpgo"
)

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		mock := httpgo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("this is a mock and we need an error"),
		}
		fmt.Printf("this is the mock %v", mock)
	})
	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		mock := httpgo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResBody:       `"current_user_url": definitellynotastring`,
			ResStatusCode: http.StatusOK,
		}
		fmt.Printf("this is the mock %v", mock)
	})
	t.Run("TestNoError", func(t *testing.T) {
		mock := httpgo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResBody:       `"current_user_url": "https://api.github.com/user"`,
			ResStatusCode: http.StatusOK,
		}
		fmt.Printf("this is the mock %v", mock)
	})
	// endpoints, err := GetEndpoints()
	// fmt.Println(err)
	// fmt.Println(endpoints)
}
