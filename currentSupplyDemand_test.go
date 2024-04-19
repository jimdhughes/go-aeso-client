package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMockResponseCurrentSupplyDemand(t *testing.T) {
	const json = `
	{
			"timestamp": "2024-04-19 03:39:27.667+0000",
			"responseCode": "200",
			"return": {
				"last_updated_datetime_utc": "2024-04-19 03:38",
				"last_updated_datetime_mpt": "2024-04-18 21:38",
				"total_max_generation_capability": 21191,
				"total_net_generation": 9932,
				"net_to_grid_generation": 7291,
				"net_actual_interchange": -276,
				"alberta_internal_load": 10208,
				"contingency_reserve_required": 445,
				"dispatched_contigency_reserve_total": 460,
				"dispatched_contingency_reserve_gen": 351,
				"dispatched_contingency_reserve_other": 109,
				"lssi_armed_dispatch": 0,
				"lssi_offered_volume": 237,
				"generation_data_list": [
					{
						"fuel_type": "COAL",
						"aggregated_maximum_capability": 820,
						"aggregated_net_generation": 396,
						"aggregated_dispatched_contingency_reserve": 0
					},
					{
						"fuel_type": "DUAL FUEL",
						"aggregated_maximum_capability": 466,
						"aggregated_net_generation": 466,
						"aggregated_dispatched_contingency_reserve": 0
					},
					{
						"fuel_type": "ENERGY STORAGE",
						"aggregated_maximum_capability": 190,
						"aggregated_net_generation": 0,
						"aggregated_dispatched_contingency_reserve": 89
					},
					{
						"fuel_type": "GAS",
						"aggregated_maximum_capability": 12246,
						"aggregated_net_generation": 6936,
						"aggregated_dispatched_contingency_reserve": 60
					},
					{
						"fuel_type": "HYDRO",
						"aggregated_maximum_capability": 894,
						"aggregated_net_generation": 108,
						"aggregated_dispatched_contingency_reserve": 202
					},
					{
						"fuel_type": "OTHER",
						"aggregated_maximum_capability": 444,
						"aggregated_net_generation": 204,
						"aggregated_dispatched_contingency_reserve": 0
					},
					{
						"fuel_type": "SOLAR",
						"aggregated_maximum_capability": 1650,
						"aggregated_net_generation": 0,
						"aggregated_dispatched_contingency_reserve": 0
					},
					{
						"fuel_type": "WIND",
						"aggregated_maximum_capability": 4481,
						"aggregated_net_generation": 1822,
						"aggregated_dispatched_contingency_reserve": 0
					}
				],
				"interchange_list": [
					{
						"path": "British Columbia",
						"actual_flow": -24
					},
					{
						"path": "Montana",
						"actual_flow": -130
					},
					{
						"path": "Saskatchewan",
						"actual_flow": -122
					}
				]
			}
	}`
	r := io.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	result, err := aesoClient.GetCurrentSupplyDemandSummary()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result.TotalMaxGenerationCapability != 21191 {
		t.Errorf("Expected: 21191, Actual: %v", result.TotalMaxGenerationCapability)
	}
	if result.TotalNetGeneration != 9932 {
		t.Errorf("Expected: 9932, Actual: %v", result.TotalNetGeneration)
	}
	if result.NetToGridGeneration != 7291 {
		t.Errorf("Expected: 7291, Actual: %v", result.NetToGridGeneration)
	}
	if result.NetActualInterchange != -276 {
		t.Errorf("Expected: -276, Actual: %v", result.NetActualInterchange)
	}
	if result.AlbertaInternalLoad != 10208 {
		t.Errorf("Expected: 10208, Actual: %v", result.AlbertaInternalLoad)
	}
	if result.ContingencyReserveRequired != 445 {
		t.Errorf("Expected: 445, Actual: %v", result.ContingencyReserveRequired)
	}
	if result.DispatchedContingencyReserveTotal != 460 {
		t.Errorf("Expected: 460, Actual: %v", result.DispatchedContingencyReserveTotal)
	}
	if result.DispatchedContingencyReserveGen != 351 {
		t.Errorf("Expected: 351, Actual: %v", result.DispatchedContingencyReserveGen)
	}
	if result.DispatchedContingencyReserveOther != 109 {
		t.Errorf("Expected: 109, Actual: %v", result.DispatchedContingencyReserveOther)
	}
	if result.LssiArmedDispatched != 0 {
		t.Errorf("Expected: 0, Actual: %v", result.LssiArmedDispatched)
	}
	if result.LssiOfferedVolume != 237 {
		t.Errorf("Expected: 237, Actual: %v", result.LssiOfferedVolume)
	}
	if len(result.GenerationDataList) != 8 {
		t.Errorf("Expected: 8, Actual: %v", len(result.GenerationDataList))
	}
	if len(result.InterchangeList) != 3 {
		t.Errorf("Expected: 3, Actual: %v", len(result.InterchangeList))
	}

}
