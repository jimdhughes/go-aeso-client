package aeso

import "encoding/json"

const AESO_API_URL_CURRENT_SUPPLY_DEMAND_SUMMARY = "https://api.aeso.ca/report/v1/csd/summary/current"

type InterchangeEntry struct {
	Path       string `json:"string"`
	ActualFlow int32  `json:"actual_flow"`
}

type GenerationDataEntry struct {
	FuelType                             string `json:"fuel_type"`
	AggregatedMaximumCapability          int32  `json:"aggregated_maximum_capability"`
	AggregatedNetGeneration              int32  `json:"aggregated_net_generation"`
	AggregatedDispatchContingencyReserve int32  `json:"aggregated_dispatch_contingency_reserve"`
}

type CurrentSupplyDemandEntry struct {
	LastUpdatedDateTimeUTC            string                `json:"last_updated_datetime_utc"`
	LastUpdatedDateTimeMPT            string                `json:"last_updated_datetime_mpt"`
	TotalMaxGenerationCapability      int32                 `json:"total_max_generation_capability"`
	TotalNetGeneration                int32                 `json:"total_net_generation"`
	NetToGridGeneration               int32                 `json:"net_to_grid_generation"`
	NetActualInterchange              int32                 `json:"net_actual_interchange"`
	AlbertaInternalLoad               int32                 `json:"alberta_internal_load"`
	ContingencyReserveRequired        int32                 `json:"contingency_reserve_required"`
	DispatchedContingencyReserveTotal int32                 `json:"dispatched_contigency_reserve_total"`
	DispatchedContingencyReserveGen   int32                 `json:"dispatched_contingency_reserve_gen"`
	DispatchedContingencyReserveOther int32                 `json:"dispatched_contingency_reserve_other"`
	LssiArmedDispatched               int32                 `json:"lssi_armed_dispatched"`
	LssiOfferedVolume                 int32                 `json:"lssi_offered_volume"`
	GenerationDataList                []GenerationDataEntry `json:"generation_data_list"`
	InterchangeList                   []InterchangeEntry    `json:"interchange_list"`
}

type CurrentSupplyDemandResponse struct {
	Timestamp    string                   `json:"timestamp"`
	ResponseCode string                   `json:"responseCode"`
	Return       CurrentSupplyDemandEntry `json:"return"`
}

func (a *AesoApiService) GetCurrentSupplyDemandSummary() (CurrentSupplyDemandEntry, error) {
	var aesoRes CurrentSupplyDemandResponse
	bytes, err := a.execute(AESO_API_URL_CURRENT_SUPPLY_DEMAND_SUMMARY)
	if err != nil {
		return CurrentSupplyDemandEntry{}, err
	}
	err = json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		return CurrentSupplyDemandEntry{}, err
	}
	return aesoRes.Return, nil
}
