package aeso

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

const AESO_API_URL_POOLPRICE = "https://api.aeso.ca/report/v1.1/price/poolPrice?startDate=%s&endDate=%s"

type AesoReportEntry struct {
	BeginDateTimeUTC        string `json:"begin_datetime_utc"`
	BeginDateTimeMPT        string `json:"begin_datetime_mpt"`
	PoolPrice               string `json:"pool_price"`
	ForecastPoolPrice       string `json:"forecast_pool_price"`
	RollingThirtyDayAverage string `json:"rolling_30day_avg"`
}

type AesoPoolResponse struct {
	Timestamp    string                     `json:"timestamp"`
	ResponseCode string                     `json:"responseCode"`
	Return       AesoPoolResponseReportPart `json:"return"`
}

type AesoPoolResponseReportPart struct {
	Report []AesoReportEntry `json:"Pool Price Report"`
}

type MappedPoolPrice struct {
	BeginDateTimeUTC        time.Time `json:"begin_datetime_utc"`
	PoolPrice               float64   `json:"pool_price"`
	ForecastPoolPrice       float64   `json:"forecast_pool_price"`
	RollingThirtyDayAverage float64   `json:"rolling_30day_avg"`
}

func (a *AesoApiService) GetPoolPrice(start, end time.Time) ([]MappedPoolPrice, error) {
	var res []MappedPoolPrice
	var aesoRes AesoPoolResponse
	sDateString := start.Format("2006-01-02")
	eDateString := end.Format("2006-01-02")
	bytes, err := a.execute(fmt.Sprintf(AESO_API_URL_POOLPRICE, sDateString, eDateString))
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		return []MappedPoolPrice{}, err
	}
	for _, entry := range aesoRes.Return.Report {
		mapped, err := mapReportValueToStruct(entry)
		if err != nil {
			return []MappedPoolPrice{}, err
		}
		res = append(res, mapped)
	}
	return res, nil
}

func mapReportValueToStruct(entry AesoReportEntry) (MappedPoolPrice, error) {
	var m MappedPoolPrice
	// extract UTC time
	fullDateString := fmt.Sprintf(entry.BeginDateTimeUTC)
	date, err := time.Parse("2006-01-02 15:04", entry.BeginDateTimeUTC)
	if err != nil {
		log.Printf("Error converting %s from Mountain to UTC\n", fullDateString)
		return m, err
	}
	//sanitize - as 0's for entries that are not available
	if entry.PoolPrice == "-" {
		entry.PoolPrice = "0"
	}
	if entry.RollingThirtyDayAverage == "-" {
		entry.RollingThirtyDayAverage = "0"
	}
	if entry.ForecastPoolPrice == "-" {
		entry.ForecastPoolPrice = "0"
	}
	price, err := strconv.ParseFloat(entry.PoolPrice, 64)
	if err != nil {
		return m, err
	}
	thirtyDayAvg, err := strconv.ParseFloat(entry.RollingThirtyDayAverage, 64)
	if err != nil {
		return m, err
	}
	forecastPoolPrice, err := strconv.ParseFloat(entry.ForecastPoolPrice, 64)
	if err != nil {
		return m, err
	}
	m = MappedPoolPrice{
		BeginDateTimeUTC:        date,
		PoolPrice:               price,
		RollingThirtyDayAverage: thirtyDayAvg,
		ForecastPoolPrice:       forecastPoolPrice,
	}
	return m, nil
}
