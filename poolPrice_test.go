package aeso

import (
	"testing"
	"time"
)

func TestMapReportValueToStruct(t *testing.T) {
	// create an AesoReport item
	report := AesoReportEntry{
		Date:         "2022-04-01 01",
		Price:        "100",
		ThirtyDayAvg: "101",
		AilDemand:    "102",
	}
	// what do we expect the value to be?
	expectedMapping := MappedPoolPrice{
		Date:         time.Date(2022, 4, 1, 7, 59, 59, 0, time.UTC),
		Price:        100,
		ThirtyDayAvg: 101,
		AILDemand:    102,
	}

	// map the report to the struct
	mappedReport, err := mapReportValueToStruct(report)
	if err != nil {
		t.Fail()
	}
	if mappedReport.Date != expectedMapping.Date {
		t.Errorf("Expected %s, got %s", expectedMapping.Date, mappedReport.Date)
	}
	if mappedReport.Price != expectedMapping.Price {
		t.Errorf("Expected %f, got %f", expectedMapping.Price, mappedReport.Price)
	}
	if mappedReport.ThirtyDayAvg != expectedMapping.ThirtyDayAvg {
		t.Errorf("Expected %f, got %f", expectedMapping.ThirtyDayAvg, mappedReport.ThirtyDayAvg)
	}
	if mappedReport.AILDemand != expectedMapping.AILDemand {
		t.Errorf("Expected %f, got %f", expectedMapping.AILDemand, mappedReport.AILDemand)
	}
}

func TestMapReportValueToStructWhenHoursAndMinutesReturned(t *testing.T) {
	// this test is to ensure that should the API change, we won't break our expected "hour represents hour ending" requirement
	report := AesoReportEntry{
		Date:         "2022-04-01 01:01",
		Price:        "100",
		ThirtyDayAvg: "101",
		AilDemand:    "102",
	}
	// what do we expect the value to be?
	expectedMapping := MappedPoolPrice{
		Date:         time.Date(2022, 4, 1, 7, 59, 59, 0, time.UTC),
		Price:        100,
		ThirtyDayAvg: 101,
		AILDemand:    102,
	}
	mappedReport, err := mapReportValueToStruct(report)
	if err != nil {
		t.Fail()
	}
	if mappedReport.Date != expectedMapping.Date {
		t.Errorf("Expected %s, got %s", expectedMapping.Date, mappedReport.Date)
	}
	if mappedReport.Price != expectedMapping.Price {
		t.Errorf("Expected %f, got %f", expectedMapping.Price, mappedReport.Price)
	}
	if mappedReport.ThirtyDayAvg != expectedMapping.ThirtyDayAvg {
		t.Errorf("Expected %f, got %f", expectedMapping.ThirtyDayAvg, mappedReport.ThirtyDayAvg)
	}
	if mappedReport.AILDemand != expectedMapping.AILDemand {
		t.Errorf("Expected %f, got %f", expectedMapping.AILDemand, mappedReport.AILDemand)
	}
}

func TestInvalidDateFromResponse(t *testing.T) {
	report := AesoReportEntry{
		Date:         "ABCIsEasyAs123 01",
		Price:        "100",
		ThirtyDayAvg: "101",
		AilDemand:    "102",
	}
	_, err := mapReportValueToStruct(report)
	if err == nil {
		t.Fail()
	}
}

func TestInvalidPriceExpect0(t *testing.T) {
	report := AesoReportEntry{
		Date:         "2022-04-01 01",
		Price:        "-",
		ThirtyDayAvg: "101",
		AilDemand:    "102",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err != nil {
		t.Error(err)
	}
	if mappedValue.Price != 0 {
		t.Fail()
	}
}

func TestInvalidEntryForThirtyDayAverageExpectError(t *testing.T) {
	report := AesoReportEntry{
		Date:         "2022-04-01 01",
		Price:        "-",
		ThirtyDayAvg: "xyz",
		AilDemand:    "102",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err == nil {
		t.Error(err)
	}
	if mappedValue.ThirtyDayAvg != 0 {
		t.Fail()
	}
}
func TestInvalidAilDemandMappingExpectError(t *testing.T) {
	report := AesoReportEntry{
		Date:         "2022-04-01 01",
		Price:        "-",
		ThirtyDayAvg: "0",
		AilDemand:    "abc",
	}
	mappedValue, err := mapReportValueToStruct(report)
	if err == nil {
		t.Error(err)
	}
	if mappedValue.ThirtyDayAvg != 0 {
		t.Fail()
	}
}
