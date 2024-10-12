package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMockOperatingReserveOfferControl(t *testing.T) {
	const json = `
	{
  "timestamp": "2024-10-12 00:02:03.706+0000",
  "responseCode": "200",
  "return": {
    "Operating Reserve Trade Merit Order": [
      {
        "begin_datetime_utc": "2024-01-01 07:00",
        "begin_datetime_mpt": "2024-01-01 00:00",
        "operating_reserve_blocks": [
          {
            "commodity": "ACTIVE",
            "product": "RR",
            "asset_ID": "ALS1",
            "volume": "10",
            "active_price": "-15.45",
            "premium_price": "0.00",
            "activation_price": "0.00",
            "offer_control": "Air Liquide Canada Inc."
          },
          {
            "commodity": "ACTIVE",
            "product": "RR",
            "asset_ID": "BIG",
            "volume": "72",
            "active_price": "-15.45",
            "premium_price": "0.00",
            "activation_price": "0.00",
            "offer_control": "TransAlta Energy Marketing Corp."
          }
        ]
      }
    ]
  }
}
	`

	r := io.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	result, err := aesoClient.GetOperatingReserveOfferControl(time.Now())
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1, got %d", len(result))
	}
	if len(result[0].OperatingReserveBlocks) != 2 {
		t.Errorf("Expected 2, got %d", len(result[0].OperatingReserveBlocks))
	}
	if result[0].OperatingReserveBlocks[0].Commodity != "ACTIVE" {
		t.Errorf("Expected ACTIVE, got %s", result[0].OperatingReserveBlocks[0].Commodity)
	}
	if result[0].OperatingReserveBlocks[0].Product != "RR" {
		t.Errorf("Expected RR, got %s", result[0].OperatingReserveBlocks[0].Product)
	}
	if result[0].OperatingReserveBlocks[0].AssetID != "ALS1" {
		t.Errorf("Expected ALS1, got %s", result[0].OperatingReserveBlocks[0].AssetID)
	}
	if result[0].OperatingReserveBlocks[0].Volume != "10" {
		t.Errorf("Expected 10, got %s", result[0].OperatingReserveBlocks[0].Volume)
	}
	if result[0].OperatingReserveBlocks[0].ActivePrice != "-15.45" {
		t.Errorf("Expected -15.45, got %s", result[0].OperatingReserveBlocks[0].ActivePrice)
	}
	if result[0].OperatingReserveBlocks[0].PremiumPrice != "0.00" {
		t.Errorf("Expected 0.00, got %s", result[0].OperatingReserveBlocks[0].PremiumPrice)
	}
	if result[0].OperatingReserveBlocks[0].ActivationPrice != "0.00" {
		t.Errorf("Expected 0.00, got %s", result[0].OperatingReserveBlocks[0].ActivationPrice)
	}
	if result[0].OperatingReserveBlocks[0].OfferControl != "Air Liquide Canada Inc." {
		t.Errorf("Expected Air Liquide Canada Inc., got %s", result[0].OperatingReserveBlocks[0].OfferControl)
	}
}
