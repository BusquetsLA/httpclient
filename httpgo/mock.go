package httpgo

type Mock struct {
	Method        string
	Url           string
	ReqBody       string
	ResBody       string
	Headers       string
	Error         error
	ResStatusCode int
}
