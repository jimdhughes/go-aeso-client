package aeso

import (
	"io"
	"log"
	"net/http"
)

const AESO_AUTH_HEADER="X-API-Key"

type AesoError struct {
	Timestamp string `json:"timestamp"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type AesoApiService struct {
	apiKey string
}

var service *AesoApiService

func (a *AesoApiService) execute(url string) []byte {
	client := &http.Client{}
	log.Printf("Getting: %s\n" , url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set(AESO_AUTH_HEADER, a.apiKey)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode >= 400 {
		log.Println("non-200 status code received")
	}
	defer res.Body.Close()
	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return buffer
}

func Init(key string) *AesoApiService {
	service = &AesoApiService{
		apiKey: key,
	}
	return service
}