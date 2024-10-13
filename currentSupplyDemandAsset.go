package aeso

import (
	"encoding/json"
	"strings"
)

const AESO_API_URL_CURRENT_SUPPLY_DEMAND_ASSET = "https://api.aeso.ca/report/v1/csd/generation/assets/current"

type AssetGenerationEntry struct {
	AssetID                      string `json:"asset"`
	FuelType                     string `json:"fuel_type"`
	SubFuelType                  string `json:"sub_fuel_type"`
	MaximumCapability            int32  `json:"maximum_capability"`
	NetGeneration                int32  `json:"net_generation"`
	DispatchedContingencyReserve int32  `json:"dispatched_contingency_reserve"`
}

type CurrentSupplyDemandAssetEntry struct {
	LastUpdatedDateTimeUTC string                 `json:"last_updated_datetime_utc"`
	LastUpdatedDateTimeMPT string                 `json:"last_updated_datetime_mpt"`
	AssetList              []AssetGenerationEntry `json:"asset_list"`
}

type CurrentSupplyDemandAssetResponse struct {
	Timestamp    string                        `json:"timestamp"`
	ResponseCode string                        `json:"responseCode"`
	Return       CurrentSupplyDemandAssetEntry `json:"return"`
}

// GetCurrentSupplyDemandAsset returns the current supply and demand for the specified asset IDs.
// If no asset IDs are specified, the function will return the current supply and demand for all assets.
// assetIds is a list of asset IDs to retrieve the current supply and demand for.(Max 20)
func (a *AesoApiService) GetCurrentSupplyDemandAsset(assetIds []string) (CurrentSupplyDemandAssetEntry, error) {
	var aesoRes CurrentSupplyDemandAssetResponse
	url := buildCurrentSupplyDemandAssetUrl(assetIds)
	bytes, err := a.execute(url)
	if err != nil {
		return CurrentSupplyDemandAssetEntry{}, err
	}
	err = json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		return CurrentSupplyDemandAssetEntry{}, err
	}
	return aesoRes.Return, nil
}

func buildCurrentSupplyDemandAssetUrl(assetIds []string) string {
	var url = AESO_API_URL_CURRENT_SUPPLY_DEMAND_ASSET
	if len(assetIds) > 0 {
		for _, assetId := range assetIds {
			if strings.Contains(url, "?") {
				url = url + "&assetIds=" + assetId
			} else {
				url = url + "?assetIds=" + assetId
			}
		}
	}
	return url
}
