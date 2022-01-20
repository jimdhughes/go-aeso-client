package aeso

import (
	"testing"
	"time"
)

const defaultFormat = "2006-01-02 15:04:05"

func TestUtilTimeZoneOffsetForDateInMDT(t *testing.T) {
	// MDT is from Second Sunday in March to first sunday in November Pick April 1 to be safe
	mdtDate := time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC)
	offest, err := GetTimezoneOffsetFromMSTForDate(mdtDate)
	if err != nil {
		t.Fail()
	}
	if offest != -21600 {
		t.Errorf("Expected %d, got %d", -25200, offest)
	}
}

func TestUtilTimeZoneOffsetForDateInMST(t *testing.T) {
	// MST is from First sunday in November to second Sunday in March pick December 1 to be safe
	mstDate := time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC)
	offest, err := GetTimezoneOffsetFromMSTForDate(mstDate)
	if err != nil {

		t.Fail()
	}
	if offest != -25200 {
		t.Errorf("Expected %d, got %d", -25200, offest)
	}
}

func TestConvertMDTStringToUTC(t *testing.T) {
	// MDT is from Second Sunday in March to first sunday in November Pick April 1 to be safe
	mdtDate := "2022-04-01 00:00:00"
	expectedValue := "2022-04-01 06:00:00"
	date, err := ConvertAesoDateToUTC(mdtDate, defaultFormat)
	if err != nil {
		t.Fail()
	}
	if date.Format(defaultFormat) != expectedValue {
		t.Errorf("Expected %s, got %s", expectedValue, date.Format(defaultFormat))
	}
}

func TestConvertFirstMDTDateStringToUTC(t *testing.T) {
	// MDT is from Second Sunday in March to first sunday in November Pick April 1 to be safe
	mdtDate := "2022-03-13 00:00:00"
	expectedValue := "2022-03-13 07:00:00"
	date, err := ConvertAesoDateToUTC(mdtDate, defaultFormat)
	if err != nil {
		t.Fail()
	}
	if date.Format(defaultFormat) != expectedValue {
		t.Errorf("Expected %s, got %s", expectedValue, date.Format(defaultFormat))
	}
}

func TestConvertMSTStringToUTC(t *testing.T) {
	// MDT is from Second Sunday in March to first sunday in November Pick April 1 to be safe
	mdtDate := "2022-12-01 00:00:00"
	expectedValue := "2022-12-01 07:00:00"
	date, err := ConvertAesoDateToUTC(mdtDate, defaultFormat)
	if err != nil {
		t.Fail()
	}
	if date.Format(defaultFormat) != expectedValue {
		t.Errorf("Expected %s, got %s", expectedValue, date.Format(defaultFormat))
	}
}
