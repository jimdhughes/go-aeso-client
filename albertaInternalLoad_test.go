package aeso

import (
	"bytes"
	"errors"
	"io"
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
	sDate = sDate.Add(-1 * 24 * time.Hour)
	_, err := aesoClient.GetAlbertaInternalLoad(sDate, eDate)
	log.Printf("Got Error: %v", err)
	if err.Error() != errMsg && err == nil {
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
	sDate = sDate.Add(-1 * 24 * time.Hour)
	_, err := aesoClient.GetAlbertaInternalLoad(sDate, eDate)
	if err.Error() != ERR_INVALID_RESPONSE_CODE && err != nil {
		t.Errorf("Expected Error: %s. Got Error: %v", ERR_INVALID_RESPONSE_CODE, err)
	}
}

func TestMappingMockedResponse(t *testing.T) {
	const json = `
	{
		"timestamp": "2024-04-19 03:06:54.730+0000",
		"responseCode": "200",
		"return": {
			"Actual Forecast Report": [
				{
					"begin_datetime_utc": "2024-04-14 06:00",
					"begin_datetime_mpt": "2024-04-14 00:00",
					"alberta_internal_load": "9131",
					"forecast_alberta_internal_load": "9123"
				}
			]
		}
	}`
	r := io.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}
	sDate := time.Now()
	eDate := time.Now()
	sDate = sDate.Add(-1 * 24 * time.Hour)
	result, err := aesoClient.GetAlbertaInternalLoad(sDate, eDate)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	if len(result) != 1 {
		t.Errorf("Expected: 1, Actual: %v", len(result))
		return
	}
	var expectedResult = result[0]
	if expectedResult.BeginDateTimeUTC.Year() != 2024 {
		t.Errorf("Expected: 2022, Actual: %v", result[0].BeginDateTimeUTC.Year())
	}
	if expectedResult.AlbertaInternalLoad != 9131 {
		t.Errorf("Expected: 9131, Actual: %v", result[0].AlbertaInternalLoad)
	}
	if expectedResult.ForecastAlbertaInternalLoad != 9123 {
		t.Errorf("Expected: 9123, Actual: %v", result[0].ForecastAlbertaInternalLoad)
	}

}

func init() {
	// initialize the aeso client for mocked responses
	aesoClient = AesoApiService{
		apiKey:     "",
		httpClient: &mocks.MockClient{},
	}
}
