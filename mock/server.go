package mock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/BusquetsLA/httpclient/core"
)

var (
	mockupServer = mockServer{ // contains all the mocks
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

type mockServer struct {
	enable     bool
	mutex      sync.Mutex
	mocks      map[string]*Mock
	httpClient core.HttpClient
}

func StartMockServer() {
	// if i didn't get that wrong, this would ensure that no matter the ammount of goroutines that reach the function
	// only one goes through and the other ones stay in the Lock(), until it's unlocked and then another one goes and so on
	mockupServer.mutex.Lock()         // so one goroutine goes past this point and the others get locked
	defer mockupServer.mutex.Unlock() // and when the function finishes it unlocks the mutex, pegarle una leida a esto igual
	mockupServer.enable = true
}

func StopMockServer() {
	// stops the mock server, when doing requests with the mock server stopped (mockupServer.enable = false)
	// the real endpoint will be called when making a request
	mockupServer.mutex.Lock()
	defer mockupServer.mutex.Unlock()
	mockupServer.enable = false
}

func AddMock(mock Mock) {
	mockupServer.mutex.Lock()
	defer mockupServer.mutex.Unlock()
	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.ReqBody)
	mockupServer.mocks[key] = &mock
}

func DeleteMock() { // empties the mockupServer of mocks
	mockupServer.mutex.Lock()
	defer mockupServer.mutex.Unlock()
	mockupServer.mocks = make(map[string]*Mock)
}

func IsMockServerEnabled() bool { // when mockupServer.enable = true the http requests made will be done to the mock server
	return mockupServer.enable
}

func GetMockedClient() core.HttpClient {
	return mockupServer.httpClient
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.clearBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil)) // O.o
	return key
}

func (m *mockServer) clearBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}
