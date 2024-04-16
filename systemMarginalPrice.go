package aeso

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const AESO_API_URL_SYSTEMMARGINALPRICE = "https://api.aeso.ca/report/v1.1/price/systemMarginalPrice?startDate=%s&endDate=%s"
const AESO_API_URL_CURRENT_SYSTEMMARGINALPRICE = "https://api.aeso.ca/report/v1.1/price/systemMarginalPrice/current"

type MappedSystemMarginalPrice struct {
	BeginDateTimeUTC    time.Time `json:"beginDateTimeUTC"`
	EndDateTimeUTC      time.Time `json:"endDateTimeUTC"`
	SystemMarginalPrice float64   `json:"systemMarginalPrice"`
	Volume              float64   `json:"volume"`
}

type AesoSystemMarginalPriceReport struct {
	BeginDateTimeUTC    string `json:"begin_datetime_utc"`
	BeginDateTimeMPT    string `json:"begin_datetime_mpt"`
	EndDateTimeUTC      string `json:"end_datetime_utc"`
	EndDateTimeMPT      string `json:"end_datetime_mpt"`
	SystemMarginalPrice string `json:"system_marginal_price"`
	Volume              string `json:"volume"`
}

type AesoSystemMarginalPriceResponseReportPart struct {
	Report []AesoSystemMarginalPriceReport `json:"System Marginal Price Report"`
}

type AesoSystemMarginalPriceResponse struct {
	Timestamp    string                                    `json:"timestamp"`
	ResponseCode string                                    `json:"responseCode"`
	Return       AesoSystemMarginalPriceResponseReportPart `json:"return"`
}

func (a *AesoApiService) GetSystemMarginalPrice(start, end time.Time) ([]MappedSystemMarginalPrice, error) {
	var response AesoSystemMarginalPriceResponse
	var res []MappedSystemMarginalPrice
	sDateString := start.Format("2006-01-02")
	eDateString := end.Format("2006-01-02")
	bytes, err := a.execute(fmt.Sprintf(AESO_API_URL_SYSTEMMARGINALPRICE, sDateString, eDateString))
	if err != nil {
		return []MappedSystemMarginalPrice{}, err
	}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return []MappedSystemMarginalPrice{}, err
	}
	for _, entry := range response.Return.Report {
		mappedValue, err := mapAesoSystemMarginalPriceToStruct(entry)
		if err != nil {
			return []MappedSystemMarginalPrice{}, err
		}
		res = append(res, mappedValue)
	}
	return res, nil
}

func (a *AesoApiService) GetCurrentSystemMarginalPrice() ([]MappedSystemMarginalPrice, error) {
	var response AesoSystemMarginalPriceResponse
	var res []MappedSystemMarginalPrice
	bytes, err := a.execute(AESO_API_URL_CURRENT_SYSTEMMARGINALPRICE)
	if err != nil {
		return []MappedSystemMarginalPrice{}, err
	}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return []MappedSystemMarginalPrice{}, err
	}
	for _, entry := range response.Return.Report {
		mappedValue, err := mapAesoSystemMarginalPriceToStruct(entry)
		if err != nil {
			return []MappedSystemMarginalPrice{}, err
		}
		res = append(res, mappedValue)
	}
	return res, nil
}

func mapAesoSystemMarginalPriceToStruct(entry AesoSystemMarginalPriceReport) (MappedSystemMarginalPrice, error) {
	var m MappedSystemMarginalPrice
	timeInUTC, err := time.Parse("2006-01-02 15:04", entry.BeginDateTimeUTC)
	if err != nil {
		return m, err
	}
	timeInUTC2, err := time.Parse("2006-01-02 15:04", entry.EndDateTimeUTC)
	if err != nil {
		return m, err
	}
	systemMarginalPrice, err := strconv.ParseFloat(entry.SystemMarginalPrice, 64)
	if err != nil {
		return m, err
	}
	volume, err := strconv.ParseFloat(entry.Volume, 64)
	if err != nil {
		return m, err
	}
	m = MappedSystemMarginalPrice{
		BeginDateTimeUTC:    timeInUTC,
		EndDateTimeUTC:      timeInUTC2,
		SystemMarginalPrice: systemMarginalPrice,
		Volume:              volume,
	}
	return m, nil
}
