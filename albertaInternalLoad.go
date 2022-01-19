package aeso

import "time"

type MappedAlbertaInternalLoad struct {
	BeginDateTimeUTC            time.Time `json:"beginDateTimeUTC"`
	BeginDateTimeMPT            time.Time `json:"beginDateTimeMPT"`
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

func (a *AesoApiService) GetAlbertaInternalLoad(start, end time.Time, offset int64) []MappedAlbertaInternalLoad {
	var res []MappedAlbertaInternalLoad
	// TODO
	return res
}
