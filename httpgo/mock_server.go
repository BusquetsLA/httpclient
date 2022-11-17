package httpgo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

var (
	server = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enable bool
	mutex  sync.Mutex
	mocks  map[string]*Mock
}

func AddMock(mock Mock) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	key := server.getMockKey(mock.Method, mock.Url, mock.ReqBody)
	server.mocks[key] = &mock
}

func StartMockServer() {
	// if i didn't get that wrong, this would ensure that no matter the ammount of goroutines that reach the function
	// only one goes through and the other ones stay in the Lock(), until it's unlocked and then another one goes and so on
	server.mutex.Lock()         // so one goroutine goes past this point and the others get locked
	defer server.mutex.Unlock() // and when the function finishes it unlocks the mutex, pegarle una leida a esto igual
	server.enable = true
}

func StopMockServer() {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.enable = false
}

func ClearMockServer() { // empties the server of mocks
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.mocks = make(map[string]*Mock)
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

func (m *mockServer) getMock(method, url, body string) *Mock {
	if !m.enable { // if there isn't a mock the library will make the call to the api
		return nil
	}

	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: fmt.Errorf("no mock matching %s from '%s' with given body", method, url),
	}
}
