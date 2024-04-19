package aeso

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/jimdhughes/go-aeso-client/mocks"
)

func TestMockPoolParticipant(t *testing.T) {
	const json = `
	{
		"timestamp": "2024-04-19 04:04:47.911+0000",
		"responseCode": "200",
		"return": [
			{
				"pool_participant_name": "1772387 Alberta Limited Partnership",
				"pool_participant_ID": "ENE2",
				"corporate_contact": "2000 10423 101 St NW \n2000 10423 101 St NW \nEdmonton, Alberta T5H 0E8 \n Phone: 780-412-3959 \n Fax: 000-000-0000 \n Attn: Lianne Redmond \n"
			},
			{
				"pool_participant_name": "1772387 Alberta Limited Partnership",
				"pool_participant_ID": "ENEP",
				"corporate_contact": "2000 10423 101 Street NW \n2000 10423 101 Street NW \nEdmonton, Alberta T5H 0E8 \n Phone: 780-412-3959 \n Fax: 000-000-0000 \n Attn: Lianne Redmond \n"
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
	result, err := aesoClient.GetPoolPriceParticpant()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected: 1, Actual: %v", len(result))
	}
	if result[0].ID != "ENE2" {
		t.Errorf("Expected: 1, Actual: %v", result[0].ID)
	}
	if result[0].Name != "1772387 Alberta Limited Partnership" {
		t.Errorf("Expected: 1, Actual: %v", result[0].Name)
	}
	if result[0].CorporateContact != "2000 10423 101 St NW \n2000 10423 101 St NW \nEdmonton, Alberta T5H 0E8 \n Phone: 780-412-3959 \n Fax: 000-000-0000 \n Attn: Lianne Redmond \n" {
		t.Errorf("Expected: 1, Actual: %v", result[0].CorporateContact)
	}
}
