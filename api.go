package aeso

import (
	"errors"
	"io"
	"log"
	"net/http"
)

const AESO_AUTH_HEADER = "X-API-Key"

type AesoError struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Details   string `json:"details"`
}

type AesoApiService struct {
	apiKey string
}

func (a *AesoApiService) execute(url string) ([]byte, error) {
	client := &http.Client{}
	log.Printf("Getting: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return []byte{}, nil
	}
	req.Header.Set(AESO_AUTH_HEADER, a.apiKey)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte{}, nil
	}
	if res.StatusCode >= 400 {
		log.Println("non-200 status code received")
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
		apiKey: key,
	}, nil
}
