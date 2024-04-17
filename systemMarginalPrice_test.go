package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

var aesoClient AesoApiService

func TestMapSystemMarginalPriceReportToStruct(t *testing.T) {
	report := AesoSystemMarginalPriceReport{
		BeginDateTimeUTC:    "2024-04-16 05:33",
		EndDateTimeUTC:      "2024-04-16 06:00",
		SystemMarginalPrice: "1.0",
		Volume:              "2.0",
	}
	expected := MappedSystemMarginalPrice{
		BeginDateTimeUTC:    time.Date(2024, 4, 16, 05, 33, 0, 0, time.UTC),
		EndDateTimeUTC:      time.Date(2024, 4, 16, 06, 00, 0, 0, time.UTC),
		SystemMarginalPrice: 1.0,
		Volume:              2.0,
	}
	result, err := mapAesoSystemMarginalPriceToStruct(report)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, result)
	}
}

func TestMockedSystemMarginalPriceCall(t *testing.T) {
	const json = `
	{
		"timestamp": "2024-04-17 03:26:44.262+0000",
		"responseCode": "200",
		"return": {
			"System Marginal Price Report": [
				{
					"begin_datetime_utc": "2024-04-16 05:33",
					"end_datetime_utc": "2024-04-16 06:00",
					"begin_datetime_mpt": "2024-04-15 23:33",
					"end_datetime_mpt": "2024-04-16 00:00",
					"system_marginal_price": "30.41",
					"volume": "50"
				}
			]
		}
	}`
	r := io.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	sDate, eDate := time.Now(), time.Now()
	sDate = sDate.Add(-1 * 24 * time.Hour)
	result, err := aesoClient.GetSystemMarginalPrice(sDate, eDate)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected: 1, Actual: %v", len(result))
	}
	if result[0].BeginDateTimeUTC.Year() != 2024 {
		t.Errorf("Expected: 2024, Actual: %v", result[0].BeginDateTimeUTC.Year())
	}
}

func init() {
	// initialize the aeso client for mocked responses
	aesoClient = AesoApiService{
		apiKey:     "",
		httpClient: &mocks.MockClient{},
	}
}
