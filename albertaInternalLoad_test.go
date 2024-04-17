package aeso

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/jimdhughes/go-aeso-client/mocks"
)
func TestMapReportValueToStructExpectSucess(t *testing.T) {
	input := AesoAlbertaInternalLoadResponseReport{
		BeginDateTimeUTC:            "2022-01-19 07:00",
		BeginDateTimeMPT:            "2022-01-19 00:00",
		AlbertaInternalLoad:         "0.0",
		ForecastAlbertaInternalLoad: "0.0",
	}
	expected := MappedAlbertaInternalLoad{
		BeginDateTimeUTC:            time.Date(2022, 1, 19, 7, 0, 0, 0, time.UTC),
		AlbertaInternalLoad:         0.0,
		ForecastAlbertaInternalLoad: 0.0,
	}
	actual, err := mapResponseToInternalLoadStruct(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestMapReportValueToStructExpectFailureOnDate(t *testing.T) {
	input := AesoAlbertaInternalLoadResponseReport{
		BeginDateTimeUTC:            "abcdefg",
		BeginDateTimeMPT:            "2022-01-19 00:00",
		AlbertaInternalLoad:         "0.0",
		ForecastAlbertaInternalLoad: "0.0",
	}
	_, err := mapResponseToInternalLoadStruct(input)
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestMapReportValueToStructExpectFailureFloatParsing(t *testing.T) {
	input := AesoAlbertaInternalLoadResponseReport{
		BeginDateTimeUTC:            "2022-01-19 07:00",
		BeginDateTimeMPT:            "2022-01-19 00:00",
		AlbertaInternalLoad:         "ab",
		ForecastAlbertaInternalLoad: "ab",
	}
	_, err := mapResponseToInternalLoadStruct(input)
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestHandleAesoResponseExpectFailure(t *testing.T) {
	const errMsg = "Error from the web server"
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(errMsg)
	}
	sDate := time.Now()
	eDate := time.Now()
	sDate.Add(-1 * 24 * time.Hour)
	_, err := aesoClient.GetAlbertaInternalLoad(sDate, eDate)
	log.Printf("Got Error: %v", err)
	if err == nil && err.Error() != errMsg {
		t.Errorf("Expected Error: %s. Expected Error: %v", errMsg, err)
	}
}

func TestHandleAesoResponseExpect400ResponseAndErr(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	sDate := time.Now()
	eDate := time.Now()
	sDate.Add(-1 * 24 * time.Hour)
	_, err := aesoClient.GetAlbertaInternalLoad(sDate, eDate)
	if err == nil && err.Error() != ERR_INVALID_RESPONSE_CODE {
		t.Errorf("Expected Error: %s. Expected Error: %v", ERR_INVALID_RESPONSE_CODE, err)
	}
}

func TestHandleAesoResponseExpectValidResponse(t *testing.T) {
	json := `{"timestamp":"2022-01-19 07:00","responseCode":"200","return":{"Actual Forecast Report":[{"begin_date_time_utc":"2022-01-19 07:00","begin_date_time_mpt":"2022-01-19 00:00","alberta_internal_load":"0.0","forecast_alberta_internal_load":"0.0"}]}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}
	sDate := time.Now()
	eDate := time.Now()
	sDate.Add(-1 * 24 * time.Hour)
	response, err := aesoClient.GetAlbertaInternalLoad(sDate, eDate)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(response) != 1 {
		t.Errorf("Expected 1 item in response, got %d", len(response))
	}
}

func init() {
	// initialize the aeso client for mocked responses
	aesoClient = AesoApiService{
		apiKey:     "",
		httpClient: &mocks.MockClient{},
	}
}
