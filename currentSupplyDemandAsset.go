package aeso

const AESO_API_URL_CURRENT_SUPPLY_DEMAND_ASSET = "https://api.aeso.ca/report/v1/csd/generation/assets/current"

type AssetGenerationEntry struct {
	AssetID                      string `json:"asset_id"`
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
