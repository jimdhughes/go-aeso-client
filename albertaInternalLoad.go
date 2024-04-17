package aeso

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const AESO_API_URL_ALBERTAINTERNALLOAD = "https://api.aeso.ca/report/v1/load/albertaInternalLoad?startDate=%s&endDate=%s"

type MappedAlbertaInternalLoad struct {
	BeginDateTimeUTC            time.Time `json:"beginDateTimeUTC"`
	AlbertaInternalLoad         float64   `json:"actualInternalLoad"`
	ForecastAlbertaInternalLoad float64   `json:"forecastInternalLoad"`
}

type AesoAlbertaInternalLoadResponseReport struct {
	BeginDateTimeUTC            string `json:"begin_date_time_utc"`
	BeginDateTimeMPT            string `json:"begin_date_time_mpt"`
	AlbertaInternalLoad         string `json:"alberta_internal_load"`
	ForecastAlbertaInternalLoad string `json:"forecast_alberta_internal_load"`
}

type AesoAlbertaInternalLoadResponseReportPart struct {
	Report []AesoAlbertaInternalLoadResponseReport `json:"Actual Forecast Report"`
}

type AesoAlbertaInternalLoadResponse struct {
	Timestamp    string                                    `json:"timestamp"`
	ResponseCode string                                    `json:"responseCode"`
	Return       AesoAlbertaInternalLoadResponseReportPart `json:"return"`
}

func (a *AesoApiService) GetAlbertaInternalLoad(start, end time.Time) ([]MappedAlbertaInternalLoad, error) {
	var res []MappedAlbertaInternalLoad
	var aesoRes AesoAlbertaInternalLoadResponse
	var sDateString = start.Format("2006-01-02")
	var eDateString = end.Format("2006-01-02")
	bytes, err := a.execute(fmt.Sprintf(AESO_API_URL_ALBERTAINTERNALLOAD, sDateString, eDateString))
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		return res, err
	}
	for _, entry := range aesoRes.Return.Report {
		mapped, err := mapResponseToInternalLoadStruct(entry)
		if err != nil {

			return []MappedAlbertaInternalLoad{}, err
		}
		res = append(res, mapped)
	}
	return res, nil
}

func mapResponseToInternalLoadStruct(response AesoAlbertaInternalLoadResponseReport) (MappedAlbertaInternalLoad, error) {
	var m MappedAlbertaInternalLoad
	// we receive a UTC date for this API, so we will use it and ignore the mountain time
	// dates come back in the format: "2018-01-01 00:00"
	timeInUTC, err := time.Parse("2006-01-02 15:04", response.BeginDateTimeUTC)
	if err != nil {
		return m, err
	}

	abInternalLoad, err := strconv.ParseFloat(response.AlbertaInternalLoad, 64)
	if err != nil {
		return m, err
	}

	abForecastInternalLoad, err := strconv.ParseFloat(response.ForecastAlbertaInternalLoad, 64)
	if err != nil {
		return m, err
	}

	m = MappedAlbertaInternalLoad{
		BeginDateTimeUTC:            timeInUTC,
		AlbertaInternalLoad:         abInternalLoad,
		ForecastAlbertaInternalLoad: abForecastInternalLoad,
	}
	return m, nil
}
