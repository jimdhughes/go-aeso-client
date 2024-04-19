package aeso

import "encoding/json"

const AESO_API_URL_POOLPARTICIPANT = "https://api.aeso.ca/report/v1/poolparticipantlist"

type ParticipantPoolEntryAgent struct {
	ID   string `json:"agent_ID"`
	Name string `json:"agent_name"`
}

type PoolParticipantEntry struct {
	ID               string                       `json:"pool_participant_ID"`
	Name             string                       `json:"pool_participant_name"`
	CorporateContact string                       `json:"corporate_contact"`
	AgentList        *[]ParticipantPoolEntryAgent `json:"agent_list"`
}

type PoolParticipantResponse struct {
	Timestamp    string                 `json:"timestamp"`
	ResponseCode string                 `json:"responseCode"`
	Return       []PoolParticipantEntry `json:"return"`
}

// GetPoolPriceParticpant retrieves the pool price participant data from the AESO API
func (a *AesoApiService) GetPoolPriceParticpant() ([]PoolParticipantEntry, error) {
	var res []PoolParticipantEntry
	var aesoRes PoolParticipantResponse
	bytes, err := a.execute(AESO_API_URL_POOLPARTICIPANT)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		return []PoolParticipantEntry{}, err
	}
	return aesoRes.Return, nil
}

// GetPoolPriceByParticipantIDsorNames retrieves the pool price participant data from the AESO API by participant IDs or names.
func (a *AesoApiService) GetPoolPriceByParticipantIDsorNames(participantIDs, participantNames []string) ([]PoolParticipantEntry, error) {
	currUrl := AESO_API_URL_POOLPARTICIPANT
	if len(participantIDs) > 0 {
		currUrl += "?participantID="
		for i, id := range participantIDs {
			if i > 0 {
				currUrl += ","
			}
			currUrl += id
		}
	}
	if len(participantNames) > 0 {
		currUrl += "?participantName="
		for i, name := range participantNames {
			if i > 0 {
				currUrl += ","
			}
			currUrl += name
		}
	}
	bytes, err := a.execute(currUrl)
	if err != nil {
		return []PoolParticipantEntry{}, err
	}
	var aesoRes PoolParticipantResponse
	err = json.Unmarshal(bytes, &aesoRes)
	if err != nil {
		return []PoolParticipantEntry{}, err
	}
	return aesoRes.Return, nil
}
