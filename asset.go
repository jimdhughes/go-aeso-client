package aeso

import (
	"encoding/json"
	"fmt"
)

const AESO_API_URL_ASSET_REPORT = "https://api.aeso.ca/report/v1/assetlist?operating_status=%s&asset_type=%s"

type AssetReportPart struct {
	AssetName            string `json:"asset_name"`
	AssetID              string `json:"asset_ID"`
	AssetType            string `json:"asset_type"`
	OperatingStatus      string `json:"operating_status"`
	PoolParticipantName  string `json:"pool_participant_name"`
	PoolParticipantID    string `json:"pool_participant_ID"`
	NetToGridAssetFlag   string `json:"net_to_grid_asset_flag"`
	AssetInclStorageFlag string `json:"asset_incl_storage_flag"`
}

type AssetReport struct {
	Timestamp    string            `json:"timestamp"`
	ResponseCode string            `json:"responseCode"`
	Return       []AssetReportPart `json:"return"`
}

func (a *AesoApiService) GetAssetReport(operatingStatus, assetType string, assetIDs []string, poolParticipantIds []string) ([]AssetReportPart, error) {
	currUrl := buildAssetUrlRequest(operatingStatus, assetType, assetIDs, poolParticipantIds)
	result, err := a.execute(currUrl)
	if err != nil {
		return []AssetReportPart{}, err
	}
	var response AssetReport
	err = json.Unmarshal(result, &response)
	if err != nil {
		return []AssetReportPart{}, err
	}
	return response.Return, nil
}

func buildAssetUrlRequest(operatingStatus, assetType string, assetIDs, poolParticipantIds []string) string {
	currUrl := fmt.Sprintf(AESO_API_URL_ASSET_REPORT, operatingStatus, assetType)
	if len(assetIDs) > 0 {
		// create a comma-separated list of asset ids to include as a parameter to the reults string
		assetIDsString := ""
		for i, assetID := range assetIDs {
			if i > 0 {
				assetIDsString += ","
			}
			assetIDsString += assetID
		}
		currUrl += "&asset_ID=" + assetIDsString
	}
	if len(poolParticipantIds) > 0 {
		// create a comma-separated list of participant ids to include as a parameter to the results tring
		poolParticipantIdsString := ""
		for i, poolParticipantId := range poolParticipantIds {
			if i > 0 {
				poolParticipantIdsString += ","
			}
			poolParticipantIdsString += poolParticipantId
		}
		currUrl += "&pool_participant_ID=" + poolParticipantIdsString
	}
	return currUrl
}
