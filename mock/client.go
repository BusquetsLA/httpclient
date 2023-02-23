package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct{}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	reqBody, err := req.GetBody()
	if err != nil {
		return nil, err
	}
	defer reqBody.Close()

	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	var res http.Response

	mock := mockupServer.mocks[mockupServer.getMockKey(req.Method, req.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		res.Body = ioutil.NopCloser(strings.NewReader(mock.ResBody))
		res.StatusCode = mock.ResStatusCode
		res.ContentLength = int64(len(mock.ResBody))

		return &res, nil
	}

	return nil, fmt.Errorf("no mock matching %s from '%s' with given body", req.Method, req.URL.String())
	// return nil, errors.New(fmt.Sprintf("no mock matching %s from '%s' with given body", req.Method, req.URL.String()))
}
