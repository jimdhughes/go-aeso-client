package aeso

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMapSystemMarginalPriceReportToStruct(t *testing.T) {
	report := AesoSystemMarginalPriceReport{
		DateHourEnding: "01/19/2022 24",
		Time:           "23:59",
		PriceInDollar:  "1.0",
		VolumeInMW:     "2.0",
	}
	expected := MappedSystemMarginalPrice{
		Date:       time.Date(2022, 1, 20, 05, 59, 0, 0, time.UTC),
		Price:      1.0,
		VolumeInMW: 2.0,
		HourEnding: 24,
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
	const json = `{"timestamp": "2022-01-20 01:53:16.840+0000","responseCode": "200","return": {"System Marginal Price Report": [{"dateHourEnding": "01/18/2022 24","time": "23:56","priceInDollar": "84.06","volumeInMW": "80"}]}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	sDate, eDate := time.Now(), time.Now()
	sDate.Add(-1 * 24 * time.Hour)
	result, err := aesoClient.GetSystemMarginalPrice(sDate, eDate)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected: 1, Actual: %v", len(result))
	}
	if result[0].Date.Year() != 2022 {
		t.Errorf("Expected: 2022, Actual: %v", result[0].Date.Year())
	}
}

func TestMapSystemMarginalPriceReportToStructWithValidResponseInvalidTime(t *testing.T) {
	report := AesoSystemMarginalPriceReport{
		DateHourEnding: "01/19/2022 24",
		Time:           "24:59",
		PriceInDollar:  "1.0",
		VolumeInMW:     "2.0",
	}
	expected := MappedSystemMarginalPrice{
		Date:       time.Date(2022, 1, 20, 06, 59, 0, 0, time.UTC),
		Price:      1.0,
		VolumeInMW: 2.0,
		HourEnding: 24,
	}
	result, err := mapAesoSystemMarginalPriceToStruct(report)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, result)
	}
}

func init() {
	// initialize the aeso client for mocked responses
	aesoClient = AesoApiService{
		apiKey:     "",
		httpClient: &mocks.MockClient{},
	}
}
