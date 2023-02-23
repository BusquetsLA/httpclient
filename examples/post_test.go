package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/BusquetsLA/httpclient/mock"
)

func TestMain(m *testing.M) {
	fmt.Println("start test cases for pkg 'examples'")
	mock.StartMockServer() // any request made to the library will be done on mock
	os.Exit(m.Run())
}

func TestCreateRepository(t *testing.T) {
	t.Run("TestErrorPostingInGithub", func(t *testing.T) {
		errorText := "this is a mock and we need to test when we get an error from github"
		mock.DeleteMock()
		mock.AddMock(mock.Mock{
			Method:  http.MethodPost,
			Url:     "https://api.github.com/user/repos",
			ReqBody: `{"name":"testing-repository","description":"","private":true}`,
			Error:   errors.New(errorText),
		})
		repository := Repository{
			Name:    "testing-repository",
			Private: true,
		}
		newRepository, err := CreateRepository(repository)
		if newRepository != nil {
			t.Error("no new repository expected when we get an error from github")
		}
		if err == nil {
			t.Errorf("error expected: %s", err.Error())
		}
		if err.Error() != errorText {
			t.Errorf(`invalid error message recieved, wanted"%s" but got "%s"`, errorText, err.Error())
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		mock.DeleteMock()
		mock.AddMock(mock.Mock{
			Method:        http.MethodPost,
			Url:           "https://api.github.com/user/repos",
			ReqBody:       `{"name":"testing-repository","private":true}`,
			ResBody:       `{"id":123,"name":"testing-repository"}`,
			ResStatusCode: http.StatusCreated,
		})
		repository := Repository{
			Name:    "testing-repository",
			Private: true,
		}
		newRepository, err := CreateRepository(repository)
		if err != nil {
			t.Errorf("no error expected, but got: %s", err.Error())
		}
		if newRepository == nil {
			t.Error("new repository expected, but got: nil")
		}
		if newRepository.Name != repository.Name {
			t.Errorf(`expected "%s" as new repository name, but github returned "%s"`, repository.Name, newRepository.Name)
		}
	})
}
