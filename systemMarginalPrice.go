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

func (a *AesoApiService) GetSystemMarginalPrice(start, end time.Time, offset int64) []MappedSystemMarginalPrice {
	var response AesoSystemMarginalPriceResponse
	var res []MappedSystemMarginalPrice
	sDateString := start.Format("2006-01-02")
	eDateString := end.Format("2006-01-02")
	bytes, err := a.execute(fmt.Sprintf(AESO_API_URL_POOLPRICE, sDateString, eDateString))
	if err != nil {
		log.Println(err)
		return []MappedSystemMarginalPrice{}
	}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range response.Return.Report {
		// Date comes back as yyyy-mm-dd HH
		parts := strings.Split(entry.DateHourEnding, " ")
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
		// time comes back as a string in military time. HH:MM
		timeParts := strings.Split(entry.Time, ":")
		hours, err := strconv.ParseInt(timeParts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		date.Add(time.Duration(hours) * time.Hour)
		minutes, err := strconv.ParseInt(timeParts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		date.Add(time.Duration(minutes) * time.Minute)
		price, err := strconv.ParseFloat(entry.PriceInDollar, 64)
		if err != nil {
			log.Fatal(err)
		}
		volume, err := strconv.ParseFloat(entry.VolumeInMW, 64)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, MappedSystemMarginalPrice{
			Date:       date,
			HourEnding: h,
			Price:      price,
			VolumeInMW: volume,
		})
	}
	return res
}
