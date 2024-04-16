package aeso

import (
	"github.com/jimdhughes/go-aeso-client/mocks"
)

// TODO: Rebuild tests since the entire API response has changed

func init() {
	// initialize the aeso client for mocked responses
	aesoClient = AesoApiService{
		apiKey:     "",
		httpClient: &mocks.MockClient{},
	}
}
