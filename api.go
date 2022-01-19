package aeso

import (
	"errors"
	"io"
	"log"
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
	log.Printf("Getting: %s\n", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return []byte{}, nil
	}
	//should we validate again that the apiKey is set?
	req.Header.Set(AESO_AUTH_HEADER, a.apiKey)
	res, err := a.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	if res.StatusCode >= 400 {
		log.Println("non-success status code received")
		return []byte{}, errors.New(ERR_INVALID_RESPONSE_CODE)
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
