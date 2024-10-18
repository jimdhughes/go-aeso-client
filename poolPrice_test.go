package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMapReportValueToStruct(t *testing.T) {
	// create an AesoReport item
	report := AesoReportEntry{
		BeginDateTimeUTC:        "2024-04-01 01:00",
		PoolPrice:               "16.36",
		RollingThirtyDayAverage: "15.85",
		ForecastPoolPrice:       "62.90",
	}
	// what do we expect the value to be?
	expectedMapping := MappedPoolPrice{
		BeginDateTimeUTC:        time.Date(2024, 4, 1, 1, 0, 0, 0, time.UTC),
		PoolPrice:               16.36,
		RollingThirtyDayAverage: 15.85,
		ForecastPoolPrice:       62.90,
	}

	// map the report to the struct
	mappedReport, err := mapReportValueToStruct(report)
	if err != nil {
		t.Fail()
	}
	if mappedReport.BeginDateTimeUTC != expectedMapping.BeginDateTimeUTC {
		t.Errorf("Expected BeginDateTimeUTC %s, got %s", expectedMapping.BeginDateTimeUTC, mappedReport.BeginDateTimeUTC)
	}
	if mappedReport.PoolPrice != expectedMapping.PoolPrice {
		t.Errorf("Expected PoolPrice %f, got %f", expectedMapping.PoolPrice, mappedReport.PoolPrice)
	}
	if mappedReport.RollingThirtyDayAverage != expectedMapping.RollingThirtyDayAverage {
		t.Errorf("Expected RollingThirtyDayAverage %f, got %f", expectedMapping.RollingThirtyDayAverage, mappedReport.RollingThirtyDayAverage)
	}
	if mappedReport.ForecastPoolPrice != expectedMapping.ForecastPoolPrice {
		t.Errorf("Expected ForecastPoolPrice %f, got %f", expectedMapping.ForecastPoolPrice, mappedReport.ForecastPoolPrice)
	}
}

func TestInvalidDateFromResponse(t *testing.T) {
	report := AesoReportEntry{
		BeginDateTimeUTC:        "ABCIsEasyAs123 01",
		PoolPrice:               "100",
		RollingThirtyDayAverage: "101",
		ForecastPoolPrice:       "102",
	}
	_, err := mapReportValueToStruct(report)
	if err == nil {
		t.Fail()
	}
}

func TestInvalidPriceExpect0(t *testing.T) {
	report := AesoReportEntry{
		BeginDateTimeUTC:        "2022-04-01 01:00",
		PoolPrice:               "-",
		RollingThirtyDayAverage: "-",
		ForecastPoolPrice:       "-",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err != nil {
		t.Error(err)
	}
	if mappedValue.PoolPrice != 0 || mappedValue.RollingThirtyDayAverage != 0 || mappedValue.ForecastPoolPrice != 0 {
		t.Errorf("Expected price:0, forecast:0, thirtyDayAvg:0, got %f, %f, %f", mappedValue.PoolPrice, mappedValue.ForecastPoolPrice, mappedValue.RollingThirtyDayAverage)
	}
}

func TestInvalidPriceExpecterror(t *testing.T) {
	report := AesoReportEntry{
		BeginDateTimeUTC:        "2022-04-01 01:00",
		PoolPrice:               "abcdefg",
		RollingThirtyDayAverage: "101",
		ForecastPoolPrice:       "102",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err == nil {
		t.Error(err)
	}
	if mappedValue.PoolPrice != 0 {
		t.Fail()
	}
}

func TestInvalidEntryForThirtyDayAverageExpectError(t *testing.T) {
	report := AesoReportEntry{
		BeginDateTimeUTC:        "04/01/2022 01",
		PoolPrice:               "-",
		RollingThirtyDayAverage: "xyz",
		ForecastPoolPrice:       "102",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err == nil {
		t.Error(err)
	}
	if mappedValue.RollingThirtyDayAverage != 0 {
		t.Fail()
	}
}

func TestInvalidAilDemandMappingExpectError(t *testing.T) {
	report := AesoReportEntry{
		BeginDateTimeUTC:        "2022-04-01 01:00",
		PoolPrice:               "-",
		RollingThirtyDayAverage: "0",
		ForecastPoolPrice:       "abc",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err == nil {
		t.Error(err)
	}
	if mappedValue.RollingThirtyDayAverage != 0 {
		t.Fail()
	}
}

func TestPoolPriceMappingForHourEnding24(t *testing.T) {
	report := AesoReportEntry{
		BeginDateTimeUTC:        "2022-04-01 23:00",
		PoolPrice:               "5",
		RollingThirtyDayAverage: "0",
		ForecastPoolPrice:       "0",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err != nil {
		t.Error(err)
	}
	if mappedValue.BeginDateTimeUTC.Hour() != 23 {
		t.Errorf("Expected %d got %d", 23, mappedValue.BeginDateTimeUTC.Hour())
	}
}

func TestParsingValidJsonResponse(t *testing.T) {
	const json = `
	{
		"timestamp": "2024-04-19 03:29:58.480+0000",
		"responseCode": "200",
		"return": {
			"Pool Price Report": [
				{
					"begin_datetime_utc": "2024-04-14 06:00",
					"begin_datetime_mpt": "2024-04-14 00:00",
					"pool_price": "28.89",
					"forecast_pool_price": "28.77",
					"rolling_30day_avg": "66.02"
				},
				{
					"begin_datetime_utc": "2024-04-14 07:00",
					"begin_datetime_mpt": "2024-04-14 01:00",
					"pool_price": "32.81",
					"forecast_pool_price": "31.35",
					"rolling_30day_avg": "66.07"
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
	result, err := aesoClient.GetPoolPrice(sDate, eDate, false)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected: 1, Actual: %v", len(result))
	}
	if result[0].BeginDateTimeUTC.Year() != 2024 {
		t.Errorf("Expected: 2024, Actual: %v", result[0].BeginDateTimeUTC.Year())
	}
	if result[0].PoolPrice != 28.89 {
		t.Errorf("Expected: 28.89, Actual: %v", result[0].PoolPrice)
	}
	if result[0].ForecastPoolPrice != 28.77 {
		t.Errorf("Expected: 28.77, Actual: %v", result[0].ForecastPoolPrice)
	}
	if result[0].RollingThirtyDayAverage != 66.02 {
		t.Errorf("Expected: 66.02, Actual: %v", result[0].RollingThirtyDayAverage)
	}

}
