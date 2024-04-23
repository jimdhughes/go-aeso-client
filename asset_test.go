package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestBuildingOfAssetTestUrlWithNoOptionalRequirements(t *testing.T) {
	operatingStatus := "operating_status"
	assetType := "asset_type"
	assetIDs := []string{}
	poolParticipantIds := []string{}
	expected := "https://api.aeso.ca/report/v1/assetlist?operating_status=operating_status&asset_type=asset_type"
	actual := buildAssetUrlRequest(operatingStatus, assetType, assetIDs, poolParticipantIds)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestBuildingOfAssetTestUrlWithAssetIDs(t *testing.T) {
	operatingStatus := "operating_status"
	assetType := "asset_type"
	assetIDs := []string{"asset_id_1", "asset_id_2"}
	poolParticipantIds := []string{}
	expected := "https://api.aeso.ca/report/v1/assetlist?operating_status=operating_status&asset_type=asset_type&asset_ID=asset_id_1,asset_id_2"
	actual := buildAssetUrlRequest(operatingStatus, assetType, assetIDs, poolParticipantIds)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestBuildingOfAssetTestUrlWithPoolParticipantIds(t *testing.T) {
	operatingStatus := "operating_status"
	assetType := "asset_type"
	assetIDs := []string{}
	poolParticipantIds := []string{"pool_participant_id_1", "pool_participant_id_2"}
	expected := "https://api.aeso.ca/report/v1/assetlist?operating_status=operating_status&asset_type=asset_type&pool_participant_ID=pool_participant_id_1,pool_participant_id_2"
	actual := buildAssetUrlRequest(operatingStatus, assetType, assetIDs, poolParticipantIds)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestBuildingOfAssetTestUrlWithAssetIDsAndPoolParticipantIds(t *testing.T) {
	operatingStatus := "operating_status"
	assetType := "asset_type"
	assetIDs := []string{"asset_id_1", "asset_id_2"}
	poolParticipantIds := []string{"pool_participant_id_1", "pool_participant_id_2"}
	expected := "https://api.aeso.ca/report/v1/assetlist?operating_status=operating_status&asset_type=asset_type&asset_ID=asset_id_1,asset_id_2&pool_participant_ID=pool_participant_id_1,pool_participant_id_2"
	actual := buildAssetUrlRequest(operatingStatus, assetType, assetIDs, poolParticipantIds)
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestResponseWithMockedApiJSONResponse(t *testing.T) {
	const json = `{
		"timestamp": "2024-04-23 18:53:14.495+0000",
		"responseCode": "200",
		"return": [
			{
				"asset_name": "asset_name_1",
				"asset_ID": "asset_id_1",
				"asset_type": "asset_type_1",
				"operating_status": "asset_status_1",
				"pool_participant_name": "asset_participant_name_1",
				"pool_participant_ID": "asset_participant_id_1",
				"net_to_grid_asset_flag": "asset_net_to_grid_flag_1",
				"asset_incl_storage_flag": "asset_incl_storage_flag_1"
			},
			{
				"asset_name": "101U 1016 SR #2",
				"asset_ID": "101U",
				"asset_type": "SINK",
				"operating_status": "Retired",
				"pool_participant_name": "Kinder Morgan Canada Inc.",
				"pool_participant_ID": "TPI",
				"net_to_grid_asset_flag": "",
				"asset_incl_storage_flag": ""
			}
		]
	}`
	r := io.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	result, err := aesoClient.GetAssetReport("operating_status", "asset_type", []string{}, []string{})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected: 2, Actual: %v", len(result))
	}
	if result[0].AssetName != "asset_name_1" {
		t.Errorf("Expected: asset_name_1, Actual: %v", result[0].AssetName)
	}
	if result[0].AssetID != "asset_id_1" {
		t.Errorf("Expected: asset_id_1, Actual: %v", result[0].AssetID)
	}
	if result[0].AssetType != "asset_type_1" {
		t.Errorf("Expected: asset_type_1, Actual: %v", result[0].AssetType)
	}
	if result[0].OperatingStatus != "asset_status_1" {
		t.Errorf("Expected: asset_status_1, Actual: %v", result[0].OperatingStatus)
	}
	if result[0].PoolParticipantName != "asset_participant_name_1" {
		t.Errorf("Expected: asset_participant_name_1, Actual: %v", result[0].PoolParticipantName)
	}
	if result[0].PoolParticipantID != "asset_participant_id_1" {
		t.Errorf("Expected: asset_participant_id_1, Actual: %v", result[0].PoolParticipantID)
	}
	if result[0].NetToGridAssetFlag != "asset_net_to_grid_flag_1" {
		t.Errorf("Expected: asset_net_to_grid_flag_1, Actual: %v", result[0].NetToGridAssetFlag)
	}
	if result[0].AssetInclStorageFlag != "asset_incl_storage_flag_1" {
		t.Errorf("Expected: asset_incl_storage_flag_1, Actual: %v", result[0].AssetInclStorageFlag)
	}

}
