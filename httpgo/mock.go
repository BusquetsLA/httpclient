package httpgo

var (
	mocks map[string]*Mock
)

type Mock struct {
	Method        string
	Url           string
	ReqBody       string
	ResBody       string
	Headers       string
	Error         error
	ResStatusCode int
}

func AddMock(mock Mock) {
	key := mock.Method + mock.Url + mock.ResBody
	mocks[key] = &mock
}
