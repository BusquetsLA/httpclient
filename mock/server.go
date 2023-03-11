package mock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/BusquetsLA/httpclient/core"
)

var (
	MockupServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	httpClient core.HttpClient

	mocks map[string]*Mock
}

func (m *mockServer) Start() {
	// this would ensure that no matter the ammount of goroutines that reach the function
	// only one goes through and the other ones stay in the Lock(), until it's unlocked and then another one goes and so on
	// so one goroutine goes past this point and the others get locked
	// and when the function finishes it unlocks the mutex, pegarle una leida a esto igual
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

func (m *mockServer) Stop() {
	// stops the mock server, when doing requests with the mock server stopped (m.enable = false)
	// the real endpoint will be called when making a request
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

func (m *mockServer) GetMockedClient() core.HttpClient {
	return m.httpClient
}

func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mocks = make(map[string]*Mock)
}

func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.clearBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
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
