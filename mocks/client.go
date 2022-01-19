// Trying my hand at mocking http calls. Example from here: https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/
package mocks

import "net/http"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// The function to mock the Do method. can be set by calling test
var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
