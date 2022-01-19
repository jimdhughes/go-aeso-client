package aeso

import "time"

func GetTimezoneOffsetFromMSTForDate(date time.Time) (int, error) {
	loc, err := time.LoadLocation("America/Edmonton")
	if err != nil {
		return 0, err
	}
	_, offset := date.In(loc).Zone()
	return offset, nil
}

func ConvertAesoDateToUTC(date string) (time.Time, error) {
	// set the golang timezone to Edmonton
	loc, err := time.LoadLocation("America/Edmonton")
	if err != nil {
		return time.Time{}, err
	}
	// parse the date string
	t, err := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	if err != nil {
		return time.Time{}, err
	}
	// convert to UTC
	return t.UTC(), nil
}
