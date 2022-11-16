package examples

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/BusquetsLA/httpclient/httpgo"
)

func TestGetEndpoints(t *testing.T) {
	httpgo.StartMockServer() // to mock any request that comes in the test
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		httpgo.AddMock(httpgo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("this is a mock and we need an error"),
		})
		endpoints, err := GetEndpoints()
		if endpoints != nil { // in this case we shouldn't get endpoints because we are forcing a bad call
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Errorf("error expected: %v", err.Error())
		}
		if err.Error() != "this is a mock and we need an error" {
			t.Error("invalid error message recieved")
		}
	})
	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		httpgo.AddMock(httpgo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResBody:       `{"current_user_url": 123}`,
			ResStatusCode: http.StatusOK,
		})
		endpoints, err := GetEndpoints()
		if endpoints != nil { // in this case we shouldn't get endpoints because we are forcing a problem when unmarshaling
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Errorf("error expected: %v", err.Error())
		}
		if !strings.Contains(err.Error(), "json: cannot unmarshal number into Go struct field") {
			// fmt.Println(err.Error()) // to print the actual error
			t.Error("invalid error message recieved")
		}

	})
	t.Run("TestNoError", func(t *testing.T) {
		httpgo.AddMock(httpgo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResBody:       `{"current_user_url": "https://api.github.com/user"}`,
			ResStatusCode: http.StatusOK,
		})
		endpoints, err := GetEndpoints()
		if err != nil {
			t.Errorf("no error expected, but got: %v", err.Error())
		}
		if endpoints == nil { // in this case we shouldn't get endpoints because we are forcing a problem when unmarshaling
			t.Error("endpoints expected, but got nil")
		}
		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
