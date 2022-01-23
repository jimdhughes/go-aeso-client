package aeso

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const AESO_API_URL_SYSTEMMARGINALPRICE = "https://api.aeso.ca/report/v1/systemMarginalPrice?startDate=%s&endDate=%s"

type MappedSystemMarginalPrice struct {
	Date       time.Time `json:"date"`
	HourEnding int64     `json:"hourEnding"`
	Price      float64   `json:"price"`
	VolumeInMW float64   `json:"volumeInMW"`
}

type AesoSystemMarginalPriceReport struct {
	DateHourEnding string `json:"dateHourEnding"`
	Time           string `json:"time"`
	PriceInDollar  string `json:"priceInDollar"`
	VolumeInMW     string `json:"volumeInMW"`
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

func mapAesoSystemMarginalPriceToStruct(entry AesoSystemMarginalPriceReport) (MappedSystemMarginalPrice, error) {
	// Date comes back as yyyy-mm-dd HH
	// But we also get a time in the format HH:MM
	parts := strings.Split(entry.DateHourEnding, " ")
	datePartString := parts[0]
	timePartsString := parts[1]
	log.Println(entry.Time[0:2])
	if entry.Time[0:2] == "24" {

		// The AESO for some reason treats hour 0 as hour 24. We need to correct this.
		entry.Time = "00" + entry.Time[2:]
	}
	fullDateString := fmt.Sprintf("%s %s:00", datePartString, entry.Time)
	date, err := ConvertAesoDateToUTC(fullDateString, "01/02/2006 15:04:05")
	if err != nil {
		return MappedSystemMarginalPrice{}, err
	}

	h, err := strconv.ParseInt(timePartsString, 10, 64)
	if err != nil {
		return MappedSystemMarginalPrice{}, err
	}

	price, err := strconv.ParseFloat(entry.PriceInDollar, 64)
	if err != nil {
		return MappedSystemMarginalPrice{}, err
	}
	volume, err := strconv.ParseFloat(entry.VolumeInMW, 64)
	if err != nil {
		return MappedSystemMarginalPrice{}, err
	}
	return MappedSystemMarginalPrice{
		Date:       date,
		HourEnding: h,
		Price:      price,
		VolumeInMW: volume,
	}, nil
}
