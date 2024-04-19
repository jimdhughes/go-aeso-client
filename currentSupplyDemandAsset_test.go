package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMockCurrentSupplyDemandAsset(t *testing.T) {
	const json = `
	{
		"timestamp": "2024-04-19 03:52:16.315+0000",
		"responseCode": "200",
		"return": {
			"last_updated_datetime_utc": "2024-04-19 03:51",
			"last_updated_datetime_mpt": "2024-04-18 21:51",
			"asset_list": [
				{
					"asset": "AFG1",
					"fuel_type": "OTHER",
					"sub_fuel_type": "",
					"maximum_capability": 131,
					"net_generation": 72,
					"dispatched_contingency_reserve": 0
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
	result, err := aesoClient.GetCurrentSupplyDemandAsset()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result.AssetList) != 1 {
		t.Errorf("Expected: 1, Actual: %v", len(result.AssetList))
	}
	if result.AssetList[0].AssetID != "AFG1" {
		t.Errorf("Expected: AFG1, Actual: %v", result.AssetList[0].AssetID)
	}

}
