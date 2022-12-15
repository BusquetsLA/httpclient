package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/BusquetsLA/httpclient/httpgo"
)

func TestCreateRepository(t *testing.T) {
	httpgo.ClearMockServer()
	httpgo.AddMock(httpgo.Mock{
		Method:  http.MethodPost,
		Url:     "https://api.github.com/user/repos",
		ReqBody: `{"name":"testing-repository","description":"","private":true}`,
		Error:   errors.New("timeout from github"),
	})

	repository := Repository{
		Name:    "testing-repository",
		Private: true,
	}

	newRepository, err := CreateRepository(repository)

	if newRepository != nil {
		t.Error("no new repository expecten when we get a timeout from github")
	}

	if err == nil {
		t.Error("error expected when we get a timeout from github")
	}

	if err.Error() != "timeout from github" {
		t.Error("invalid error message")
	}

	fmt.Println(err)
	fmt.Println(newRepository)
}
