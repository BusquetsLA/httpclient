package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/BusquetsLA/httpclient/httpgo"
)

func TestCreateRepository(t *testing.T) {
	t.Run("TestErrorPostingInGithub", func(t *testing.T) {
		errorText := "this is a mock and we need to test when we get an error from github"
		httpgo.ClearMockServer()
		httpgo.AddMock(httpgo.Mock{
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
		httpgo.ClearMockServer()
		httpgo.AddMock(httpgo.Mock{
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
