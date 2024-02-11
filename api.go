package aeso

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const AESO_AUTH_HEADER = "X-API-Key"
const ERR_INVALID_RESPONSE_CODE = "invalid response code received"

type AesoError struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Details   string `json:"details"`
}

type AesoApiService struct {
	apiKey     string
	httpClient HTTPClient
}

// HTTPClient as an interface to allow mocking
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (a *AesoApiService) execute(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, nil
	}
	//should we validate again that the apiKey is set?
	req.Header.Set(AESO_AUTH_HEADER, a.apiKey)
	res, err := a.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if res.StatusCode >= 400 {
		return []byte{}, errors.New(fmt.Sprintf("%s: %d", ERR_INVALID_RESPONSE_CODE, res.StatusCode))
	}
	defer res.Body.Close()
	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, nil
	}
	return buffer, nil
}

func NewAesoApiService(key string) (AesoApiService, error) {
	if key == "" {
		return AesoApiService{}, errors.New("AESO API key is required")
	}
	return AesoApiService{
		apiKey:     key,
		httpClient: &http.Client{},
	}, nil
}
