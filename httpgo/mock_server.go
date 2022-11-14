package httpgo

import "sync"

var (
	server = mockServer{}
)

type mockServer struct {
	enable bool
	mutex  sync.Mutex // ensures a single gorutine access the mutex
	mocks  map[string]*Mock
}

func AddMock(mock Mock) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	key := mock.Method + mock.Url + mock.ResBody
	server.mocks[key] = &mock
}

func StartMockServer() {
	server.mutex.Lock()         // to ensure only one goroutine reach the lock
	defer server.mutex.Unlock() // when the function finishes it unlocks the mutex, pegarle una leida a esto igual
	server.enable = true
}
func StopMockServer() {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.enable = false
}
