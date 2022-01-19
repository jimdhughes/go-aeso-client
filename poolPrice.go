package aeso

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const AESO_API_URL_POOLPRICE = "https://api.aeso.ca/report/v1/poolPrice?startDate=%s&endDate=%s"

type AesoReportEntry struct {
	Date         string `json:"dateHourEnding"`
	Price        string `json:"priceInDollar"`
	ThirtyDayAvg string `json:"averagePoolPrice"`
	AilDemand    string `json:"ailDemand"`
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
	Date         time.Time `json:"date"`
	Price        float64   `json:"price"`
	ThirtyDayAvg float64   `json:"thirtyDayAvg"`
	AILDemand    float64   `json:"ailDemand"`
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
		log.Println(err)
		return []MappedPoolPrice{}, err
	}
	for _, entry := range aesoRes.Return.Report {
		mapped, err := mapReportValueToStruct(entry)
		if err != nil {
			log.Println(err)
			return []MappedPoolPrice{}, err
		}
		res = append(res, mapped)
	}
	return res, nil
}

func mapReportValueToStruct(entry AesoReportEntry) (MappedPoolPrice, error) {
	var m MappedPoolPrice
	// Date comes back as yyyy-mm-dd HH where HH is the hour ending (so HH:59)
	parts := strings.Split(entry.Date, " ")
	datePartString := parts[0]
	timePartsString := parts[1]
	if len(timePartsString) > 2 {
		timePartsString = timePartsString[0:2] // we expect to only get the hour back in this API call
	}
	fullDateString := fmt.Sprintf("%s %s:59:59", datePartString, timePartsString)
	date, err := ConvertAesoDateToUTC(fullDateString)
	if err != nil {
		log.Println(err)
		log.Printf("Error converting %s from Mountain to UTC\n", fullDateString)
		return m, err
	}
	if entry.Price == "-" {
		log.Println("Processed reproted prices. Setting price to be 0")
		entry.Price = "0"
	}
	price, err := strconv.ParseFloat(entry.Price, 64)
	if err != nil {
		log.Println(err)
		return m, err
	}
	thirtyDayAvg, err := strconv.ParseFloat(entry.ThirtyDayAvg, 64)
	if err != nil {
		thirtyDayAvg = 0
		return m, err
	}
	ailDemand, err := strconv.ParseFloat(entry.AilDemand, 64)
	if err != nil {
		return m, err
	}
	m = MappedPoolPrice{
		Date:         date,
		Price:        price,
		ThirtyDayAvg: thirtyDayAvg,
		AILDemand:    ailDemand,
	}
	return m, nil
}
