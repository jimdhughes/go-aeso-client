package aeso

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const AESO_API_URL_POOLPRICE="https://api.aeso.ca/report/v1/poolPrice?startDate=%s&endDate=%s"

type AesoReportEntry struct {
	Date string `json:"dateHourEnding"`
	Price string `json:"priceInDollar"`
	ThirtyDayAvg string `json:"averagePoolPrice"`
	AilDemand string `json:"ailDemand"`
}

type AesoPoolResponse struct {
	Timestamp string `json:"timestamp"`
	ResponseCode string `json:"responseCode"`
	Return AesoPoolResponseReportPart `json:"return"`
}

type AesoPoolResponseReportPart struct {
	Report []AesoReportEntry `json:"Pool Price Report"`
}

type MappedPoolPrice struct {
	Date time.Time `json:"date"`
	Price float64 `json:"price"`
	ThirtyDayAvg float64 `json:"thirtyDayAvg"`
	AILDemand float64 `json:"ailDemand"`
}

func (a *AesoApiService) GetPoolPrice(start, end time.Time, offset int64) []MappedPoolPrice {
	var res []MappedPoolPrice
	var aesoRes AesoPoolResponse
	sDateString := start.Format("2006-01-02")
	eDateString := end.Format("2006-01-02")
	bytes := a.execute(fmt.Sprintf(AESO_API_URL_POOLPRICE, sDateString, eDateString))
	err := json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(aesoRes)
	for _, entry := range(aesoRes.Return.Report) {
		// Date comes back as yyyy-mm-dd HH
		parts := strings.Split(entry.Date, " ")
		date, err := time.Parse("01/02/2006", parts[0])
		if err != nil {
			log.Fatal(err)
		}
		if len(parts[1]) > 2 {
			parts[1] = parts[1][0:1]
		}
		h, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		// Add the offset and the hour
		log.Printf("Offset: %d\n", offset)
		date = date.Add(time.Duration(h)*time.Hour).Add(time.Duration(-1*offset) * time.Second) // For each hour of the day
		if entry.Price == "-" {
			log.Println("Completed processing reported prices")
			break
		}
		price, err := strconv.ParseFloat(entry.Price, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		thirtyDayAvg, err := strconv.ParseFloat(entry.ThirtyDayAvg, 64)
		if err != nil {
			thirtyDayAvg = 0
			continue
		}
		ailDemand, err := strconv.ParseFloat(entry.AilDemand, 64)
		if err != nil {
			continue
		}
		m := MappedPoolPrice {
			Date: date,
			Price: price,
			ThirtyDayAvg: thirtyDayAvg,
			AILDemand: ailDemand,
		}
		res = append(res, m)
	}
	return res
}
