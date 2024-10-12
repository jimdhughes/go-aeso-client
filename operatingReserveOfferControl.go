package aeso

import (
	"encoding/json"
	"fmt"
	"time"
)

const AESO_API_URL_OPERATING_RESERVE_TRADE_MERIT_ORDER = "https://api.aeso.ca/report/v1/operatingReserveOfferControl?startDate=%s"

type OperatingReserveBlocks struct {
	Commodity       string `json:"commodity"`
	Product         string `json:"product"`
	AssetID         string `json:"asset_ID"`
	Volume          string `json:"volume"`
	ActivePrice     string `json:"active_price"`
	PremiumPrice    string `json:"premium_price"`
	ActivationPrice string `json:"activation_price"`
	OfferControl    string `json:"offer_control"`
}

type OperatingReserveTradeMeritOrder struct {
	BeginDateTimeUTC       string                   `json:"begin_datetime_utc"`
	BeginDateTimeMPT       string                   `json:"begin_datetime_mpt"`
	OperatingReserveBlocks []OperatingReserveBlocks `json:"operating_reserve_blocks"`
}

type OperatingReserveOfferControlReport struct {
	Timestamp    string                            `json:"timestamp"`
	ResponseCode string                            `json:"responseCode"`
	Return       []OperatingReserveTradeMeritOrder `json:"return"`
}

// GetOperatingReserveOfferControl returns a list of operating reserve trade merit orders based on the parameters provided
// start: the start date of the report. valid dates >= 2012-10-04
func (a *AesoApiService) GetOperatingReserveOfferControl(start time.Time) ([]OperatingReserveTradeMeritOrder, error) {
	sDateString := start.Format("2006-01-02")
	currUrl := buildOperatingReserveOfferControlUrlRequest(sDateString)
	result, err := a.execute(currUrl)
	if err != nil {
		return []OperatingReserveTradeMeritOrder{}, err
	}
	var response OperatingReserveOfferControlReport
	err = json.Unmarshal(result, &response)
	if err != nil {
		return []OperatingReserveTradeMeritOrder{}, err
	}
	return response.Return, nil
}

func buildOperatingReserveOfferControlUrlRequest(startDate string) string {
	currUrl := fmt.Sprintf(AESO_API_URL_OPERATING_RESERVE_TRADE_MERIT_ORDER, startDate)
	return currUrl
}
