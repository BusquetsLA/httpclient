package httpgo

import "sync"

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

func (m *mockServer) getMockKey(method, url, body string) string {
	return method + url + body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	return m.mocks[m.getMockKey(method, url, body)]
}
