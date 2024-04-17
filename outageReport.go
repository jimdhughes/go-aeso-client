package aeso

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	AESO_API_URL_OUTAGE_REPORT             = "https://api.aeso.ca/intertie/v1/outage?startDate=%s&endDate=%s"
	AESO_API_URL_OUTAGE_REPORT_BY_INTERTIE = "https://api.aeso.ca/intertie/v1/outage?startDate=%s&endDate=%s&affectedIntertieOrFlowgate=%s"
	AESO_FLOWGATE_INTERTIE_BC              = "BC"
	AESO_FLOWGATE_INTERTIE_MATL            = "MATL"
	AESO_FLOWGATE_INTERTIE_SK              = "SK"
	AESO_FLOWGATE_INTERTIE_BC_MATL         = "BC_MATL"
)

type OutageEntry struct {
	Element                    string   `json:"element"`
	AffectedIntertieOrFlowgate []string `json:"affectedIntertieOrFlowgate"`
	ToInLocalTime              string   `json:"toInLocalTime"`
	FromInLocalTime            string   `json:"fromInLocalTime"`
}
type OutagesEntry struct {
	Outage []OutageEntry `json:"Outage"`
}

type OutageReportResponse struct {
	Outages OutagesEntry `json:"Outages"`
}

func (a *AesoApiService) GetOutageReportByDate(startDate, endDate time.Time) ([]OutageEntry, error) {
	sDateString := startDate.Format("2006-01-02")
	eDateString := endDate.Format("2006-01-02")
	result, err := executeOutageReportRequest(fmt.Sprintf(AESO_API_URL_OUTAGE_REPORT, sDateString, eDateString), a)
	if err != nil {
		return []OutageEntry{}, err
	}
	return result, nil
}

func (a *AesoApiService) GetOutageReportByDateAndIntertie(startDate, endDate time.Time, intertie string) ([]OutageEntry, error) {
	sDateString := startDate.Format("2006-01-02")
	eDateString := endDate.Format("2006-01-02")
	if !isIntertieOrFlowgateValid(intertie) {
		return []OutageEntry{}, fmt.Errorf("invalid intertie or flowgate value: %s", intertie)
	}
	result, err := executeOutageReportRequest(fmt.Sprintf(AESO_API_URL_OUTAGE_REPORT_BY_INTERTIE, sDateString, eDateString, intertie), a)
	if err != nil {
		return []OutageEntry{}, err
	}
	return result, nil
}

func executeOutageReportRequest(url string, a *AesoApiService) ([]OutageEntry, error) {
	bytes, err := a.execute(url)
	if err != nil {
		return []OutageEntry{}, err
	}
	var response OutageReportResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return []OutageEntry{}, err
	}
	return response.Outages.Outage, nil
}

// There's probably a better way to do this
func isIntertieOrFlowgateValid(intertie string) bool {
	switch intertie {
	case AESO_FLOWGATE_INTERTIE_BC, AESO_FLOWGATE_INTERTIE_MATL, AESO_FLOWGATE_INTERTIE_SK, AESO_FLOWGATE_INTERTIE_BC_MATL:
		return true
	default:
		return false
	}
}
