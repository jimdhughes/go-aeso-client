package aeso

import (
	"encoding/json"
	"net/http"
	"time"
)

const AESO_API_URL_GENERATION = "https://www.aeso.ca/ets/ets-generation.json"

type TimeSeriesMeasurement struct {
	Measurement string
	Source      string
	Value       int
	Timestamp   int
}

type Measurement struct {
	Date    time.Time `json:"date"`
	Measure int       `json:"measure"`
}

type Source struct {
	Coal  []Measurement `json:"coal"`
	Gas   []Measurement `json:"gas"`
	Hydro []Measurement `json:"hydro"`
	Wind  []Measurement `json:"wind"`
	Other []Measurement `json:"other"`
}

type GenerationInfo struct {
	Updated                      string `json:"updated"`
	MaxCapacity                  Source `json:"maxCapacity"`
	DispatchedContingencyReserve Source `json:"dispatchedContingencyReserve"`
	TotalNetGeneration           Source `json:"totalNetGeneration"`
}

type AesoSource struct {
	Coal  [][]int `json:"COAL"`
	Gas   [][]int `json:"GAS"`
	Hydro [][]int `json:"HYDRO"`
	Wind  [][]int `json:"WIND"`
	Other [][]int `json:"OTHER"`
}

type AesoResponse struct {
	Updated                      string     `json:"update"`
	MaxCapacity                  AesoSource `json:"mc"`
	DispatchedContingencyReserve AesoSource `json:"dcr"`
	TotalNetGeneration           AesoSource `json:"tng"`
}

func (a *AesoApiService) GetGenerationData() (GenerationInfo, error) {
	aesoRes, err := getData()
	if err != nil {
		return GenerationInfo{}, nil
	}
	mappedResponse, err := mapAesoData(aesoRes)
	if err != nil {
		return GenerationInfo{}, nil
	}
	return mappedResponse, nil
}

func getData() (AesoResponse, error) {
	client := http.Client{}
	resp, err := client.Get(AESO_API_URL_GENERATION)
	if err != nil {
		return AesoResponse{}, err
	}
	defer resp.Body.Close()
	var data AesoResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return AesoResponse{}, err
	}
	return data, nil
}

// mapAesoData takes the AesoResponse from the API and maps it to a GenerationInfo struct
func mapAesoData(data AesoResponse) (GenerationInfo, error) {
	g := GenerationInfo{
		Updated:                      data.Updated,
		MaxCapacity:                  Source{},
		DispatchedContingencyReserve: Source{},
		TotalNetGeneration:           Source{},
	}

	// MC
	g.MaxCapacity.Coal = extractAESOMeasurements(data.MaxCapacity.Coal)
	g.MaxCapacity.Gas = extractAESOMeasurements(data.MaxCapacity.Gas)
	g.MaxCapacity.Hydro = extractAESOMeasurements(data.MaxCapacity.Hydro)
	g.MaxCapacity.Wind = extractAESOMeasurements(data.MaxCapacity.Wind)
	g.MaxCapacity.Other = extractAESOMeasurements(data.MaxCapacity.Other)
	// DCR
	g.DispatchedContingencyReserve.Coal = extractAESOMeasurements(data.DispatchedContingencyReserve.Coal)
	g.DispatchedContingencyReserve.Gas = extractAESOMeasurements(data.DispatchedContingencyReserve.Gas)
	g.DispatchedContingencyReserve.Hydro = extractAESOMeasurements(data.DispatchedContingencyReserve.Hydro)
	g.DispatchedContingencyReserve.Wind = extractAESOMeasurements(data.DispatchedContingencyReserve.Wind)
	g.DispatchedContingencyReserve.Other = extractAESOMeasurements(data.DispatchedContingencyReserve.Other)
	// TNG
	g.TotalNetGeneration.Coal = extractAESOMeasurements(data.TotalNetGeneration.Coal)
	g.TotalNetGeneration.Gas = extractAESOMeasurements(data.TotalNetGeneration.Gas)
	g.TotalNetGeneration.Hydro = extractAESOMeasurements(data.TotalNetGeneration.Hydro)
	g.TotalNetGeneration.Wind = extractAESOMeasurements(data.TotalNetGeneration.Wind)
	g.TotalNetGeneration.Other = extractAESOMeasurements(data.TotalNetGeneration.Other)
	return g, nil
}

// extractAESOMeasurements takes the expected array of [][]int and returns an array of Measurements for easier debugging
// Integer array has the structure [[unixTimestamp, value]] where unixTimestamp is the date/time and measurement is MWH produced
func extractAESOMeasurements(input [][]int) []Measurement {
	var measurements []Measurement
	for _, entry := range input {
		measurement := Measurement{
			Date:    time.Unix(int64(entry[0]/1000), 0),
			Measure: entry[1],
		}
		measurements = append(measurements, measurement)
	}
	return measurements
}
