package aeso

import (
	"encoding/json"
	"fmt"
	"time"
)

const AESO_API_URL_METERED_VOLUME = "https://api.aeso.ca/report/v1/meteredvolume/details?startDate=%s"

type MeteredVolumeItem struct {
	BeginDateUTC  string `json:"begin_date_utc"`
	BeginDateMPT  string `json:"begin_date_mpt"`
	MeteredVolume string `json:"metered_volume"`
}
type AssetItem struct {
	AssetID           string              `json:"asset_ID"`
	AssetClass        string              `json:"asset_class"`
	MeteredVolumeList []MeteredVolumeItem `json:"metered_volume_list"`
}
type MeteredVolumeReportPart struct {
	PoolParticipantID string      `json:"pool_participant_ID"`
	AssetList         []AssetItem `json:"asset_list"`
}
type MeteredVolumeReport struct {
	Timestamp    string                    `json:"timestamp"`
	ResponseCode string                    `json:"responseCode"`
	Return       []MeteredVolumeReportPart `json:"return"`
}

// MeteredVolumeReportPart represents a part of the metered volume report
// Note: If you do not provide limiting AssetIDs, or Participant IDs this is a long call
// and may take a long time to return
// startDate - the start date of the report. valid dates >= 2000-01-01
// endDate - the end date of the report. valid dates >= 2000-01-01
// assetIds - a list of asset IDs to limit the report to or empty array
// participantIds - a list of participant IDs to limit the report to or empty array
func (a *AesoApiService) GetMeteredVolume(startDate time.Time, endDate *time.Time, assetIds, participantIds []string) ([]MeteredVolumeReportPart, error) {
	sDate := startDate.Format("2006-01-02")
	eDate := ""
	if endDate != nil {
		eDate = endDate.Format("2006-01-02")
	}
	currUrl := buildMeteredVolumeUrlRequest(sDate, eDate, assetIds, participantIds)
	result, err := a.execute(currUrl)
	if err != nil {
		return []MeteredVolumeReportPart{}, err
	}
	var response MeteredVolumeReport
	err = json.Unmarshal(result, &response)
	if err != nil {
		return []MeteredVolumeReportPart{}, err
	}
	return response.Return, nil
}

func buildMeteredVolumeUrlRequest(startDate, endDate string, assetIds, participantIds []string) string {
	currUrl := fmt.Sprintf(AESO_API_URL_METERED_VOLUME, startDate)
	if len(assetIds) > 0 {
		// create a comma-separated list of asset ids to include as a parameter to the reults string
		assetIDsString := ""
		for i, assetID := range assetIds {
			if i > 0 {
				assetIDsString += ","
			}
			assetIDsString += assetID
		}
		currUrl += "&asset_ID=" + assetIDsString
	}
	if len(participantIds) > 0 {
		// create a comma-separated list of participant ids to include as a parameter to the reults string
		participantIDsString := ""
		for i, participantID := range participantIds {
			if i > 0 {
				participantIDsString += ","
			}
			participantIDsString += participantID
		}
		currUrl += "&pool_participant_ID=" + participantIDsString
	}
	if endDate != "" {
		currUrl += "&endDate=" + endDate
	}
	return currUrl
}
