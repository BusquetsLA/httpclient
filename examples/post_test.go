package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/BusquetsLA/httpclient/mock"
)

func TestCreateRepository(t *testing.T) {
	t.Run("TestErrorPostingOnGithub", func(t *testing.T) {
		errorText := "this is a mock and we need to test when we get an error from github"
		mock.MockupServer.DeleteMocks()
		mock.MockupServer.AddMock(mock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"testing-repository","private":true}`,
			Error:       errors.New(errorText),
		})
		repository := Repository{
			Name:    "testing-repository",
			Private: true,
		}
		createdRepository, err := CreateRepository(repository)
		if createdRepository != nil {
			t.Error("no new repository expected when we get an error from github")
		}
		if err == nil {
			t.Errorf("error expected: %s", err.Error())
		}
		if err.Error() != errorText {
			t.Errorf(`invalid error message recieved, wanted "%s" but got "%s"`, errorText, err.Error())
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		mock.MockupServer.DeleteMocks()
		mock.MockupServer.AddMock(mock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"testing-repository","private":true}`,
			ResponseBody:       `{"id":123,"name":"testing-repository"}`,
			ResponseStatusCode: http.StatusCreated,
		})
		repository := Repository{
			Name:    "testing-repository",
			Private: true,
		}
		createdRepository, err := CreateRepository(repository)
		if err != nil {
			t.Errorf("no error expected, but got: %s", err.Error())
		}
		if createdRepository == nil {
			t.Error("new repository expected, but got: nil")
		}
		if createdRepository.Name != repository.Name {
			t.Errorf(`expected "%s" as new repository name, but github returned "%s"`, repository.Name, createdRepository.Name)
		}
	})
}
