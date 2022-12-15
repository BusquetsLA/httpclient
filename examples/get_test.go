package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/BusquetsLA/httpclient/httpgo"
)

func TestMain(m *testing.M) {
	fmt.Println("start test cases for pkg 'examples'")
	httpgo.StartMockServer() // any request made to the library will be done on mock
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		errorText := "this is a mock and we need to test when we get an error from github"
		httpgo.ClearMockServer()
		httpgo.AddMock(httpgo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New(errorText),
		})
		endpoints, err := GetEndpoints()
		if endpoints != nil { // in this case we shouldn't get endpoints because we are forcing a bad call
			t.Error("no endpoints expected when we get an error from github")
		}
		if err == nil {
			t.Error("forcerd error expected")
		}
		if err.Error() != errorText {
			// fmt.Println(err.Error())
			t.Errorf(`invalid error message recieved, wanted "%s" but got "%s"`, errorText, err.Error())
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		errorText := "json: cannot unmarshal number into Go struct field"
		httpgo.ClearMockServer()
		httpgo.AddMock(httpgo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResBody:       `{"current_user_url": 123}`,
			ResStatusCode: http.StatusOK,
		})
		endpoints, err := GetEndpoints()
		if endpoints != nil { // in this case we shouldn't get endpoints because we are forcing a problem when unmarshaling
			t.Error("no endpoints expected when there's an error unmarshaling the response body")
		}
		if err == nil {
			t.Error("forced error expected")
		}
		if !strings.Contains(err.Error(), errorText) {
			// fmt.Println(err.Error()) // to print the actual error
			t.Errorf(`invalid error message recieved, wanted "%s" but got "%s"`, errorText, err.Error())
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		currentUserUrl := "https://api.github.com/user"
		httpgo.ClearMockServer()
		httpgo.AddMock(httpgo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResBody:       `{"current_user_url": "https://api.github.com/user"}`,
			ResStatusCode: http.StatusOK,
		})
		endpoints, err := GetEndpoints()
		if err != nil {
			t.Errorf(`no error expected, but got "%s"`, err.Error())
		}
		if endpoints == nil {
			t.Error(`endpoints expected, but got "<nil>"`)
		}
		if endpoints.CurrentUserUrl != currentUserUrl {
			t.Errorf(`invalid current user url, wanted "%s" but got "%s"`, currentUserUrl, endpoints.CurrentUserUrl)
		}
	})
}
