package aeso

import (
	"testing"
	"time"
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
