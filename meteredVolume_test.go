package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMockMeteredVolume(t *testing.T) {
	const json = `
	{
  "timestamp": "2024-10-12 17:27:06.968+0000",
  "responseCode": "200",
  "return": [
    {
      "pool_participant_ID": "EWSI",
      "asset_list": [
        {
          "asset_ID": "KKP1",
          "asset_class": "IPP",
          "metered_volume_list": [
            {
              "begin_date_utc": "2024-10-10 06:00",
              "begin_date_mpt": "2024-10-10 00:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 07:00",
              "begin_date_mpt": "2024-10-10 01:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 08:00",
              "begin_date_mpt": "2024-10-10 02:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 09:00",
              "begin_date_mpt": "2024-10-10 03:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 10:00",
              "begin_date_mpt": "2024-10-10 04:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 11:00",
              "begin_date_mpt": "2024-10-10 05:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 12:00",
              "begin_date_mpt": "2024-10-10 06:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 13:00",
              "begin_date_mpt": "2024-10-10 07:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 14:00",
              "begin_date_mpt": "2024-10-10 08:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 15:00",
              "begin_date_mpt": "2024-10-10 09:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 16:00",
              "begin_date_mpt": "2024-10-10 10:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 17:00",
              "begin_date_mpt": "2024-10-10 11:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 18:00",
              "begin_date_mpt": "2024-10-10 12:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 19:00",
              "begin_date_mpt": "2024-10-10 13:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 20:00",
              "begin_date_mpt": "2024-10-10 14:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 21:00",
              "begin_date_mpt": "2024-10-10 15:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 22:00",
              "begin_date_mpt": "2024-10-10 16:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 23:00",
              "begin_date_mpt": "2024-10-10 17:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 00:00",
              "begin_date_mpt": "2024-10-10 18:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 01:00",
              "begin_date_mpt": "2024-10-10 19:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 02:00",
              "begin_date_mpt": "2024-10-10 20:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 03:00",
              "begin_date_mpt": "2024-10-10 21:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 04:00",
              "begin_date_mpt": "2024-10-10 22:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 05:00",
              "begin_date_mpt": "2024-10-10 23:00",
              "metered_volume": "0"
            }
          ]
        },
        {
          "asset_ID": "KKP2",
          "asset_class": "IPP",
          "metered_volume_list": [
            {
              "begin_date_utc": "2024-10-10 06:00",
              "begin_date_mpt": "2024-10-10 00:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 07:00",
              "begin_date_mpt": "2024-10-10 01:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 08:00",
              "begin_date_mpt": "2024-10-10 02:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 09:00",
              "begin_date_mpt": "2024-10-10 03:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 10:00",
              "begin_date_mpt": "2024-10-10 04:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 11:00",
              "begin_date_mpt": "2024-10-10 05:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 12:00",
              "begin_date_mpt": "2024-10-10 06:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 13:00",
              "begin_date_mpt": "2024-10-10 07:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 14:00",
              "begin_date_mpt": "2024-10-10 08:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 15:00",
              "begin_date_mpt": "2024-10-10 09:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 16:00",
              "begin_date_mpt": "2024-10-10 10:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 17:00",
              "begin_date_mpt": "2024-10-10 11:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 18:00",
              "begin_date_mpt": "2024-10-10 12:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 19:00",
              "begin_date_mpt": "2024-10-10 13:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 20:00",
              "begin_date_mpt": "2024-10-10 14:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 21:00",
              "begin_date_mpt": "2024-10-10 15:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 22:00",
              "begin_date_mpt": "2024-10-10 16:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-10 23:00",
              "begin_date_mpt": "2024-10-10 17:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 00:00",
              "begin_date_mpt": "2024-10-10 18:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 01:00",
              "begin_date_mpt": "2024-10-10 19:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 02:00",
              "begin_date_mpt": "2024-10-10 20:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 03:00",
              "begin_date_mpt": "2024-10-10 21:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 04:00",
              "begin_date_mpt": "2024-10-10 22:00",
              "metered_volume": "0"
            },
            {
              "begin_date_utc": "2024-10-11 05:00",
              "begin_date_mpt": "2024-10-10 23:00",
              "metered_volume": "0"
            }
          ]
        }
      ]
    }
  ]
}
	`

	r := io.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	result, err := aesoClient.GetMeteredVolume(time.Now(), nil, []string{}, []string{})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 2, got %d", len(result))
	}
	if len(result[0].AssetList) != 2 {
		t.Errorf("Expected 2, got %d", len(result[0].AssetList))
	}
	// test the mapping of the first asset
	if result[0].AssetList[0].AssetID != "KKP1" {
		t.Errorf("Expected KKP1, got %s", result[0].AssetList[0].AssetID)
	}
	if len(result[0].AssetList[0].MeteredVolumeList) != 24 {
		t.Errorf("Expected 24, got %d", len(result[0].AssetList[0].MeteredVolumeList))
	}
	if result[0].AssetList[0].MeteredVolumeList[0].MeteredVolume != "0" {
		t.Errorf("Expected 0, got %s", result[0].AssetList[0].MeteredVolumeList[0].MeteredVolume)
	}
}
