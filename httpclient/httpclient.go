package httpclient

import "net/http"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClientMock struct {
	ResponseData *http.Response
	ErrorData    error
}

func (client HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	return client.ResponseData, client.ErrorData
}
