package dto

import "testing"

func TestGetCidadeRequest(t *testing.T) {
	response := &ViaCEPResponse{
		Localidade: "Erechim",
	}
	if response.Localidade != "Erechim" {
		t.Errorf("Expected Localidade to be 'Erechim', got '%s'", response.Localidade)
	}
	response2 := &WeatherResponse{
		Current: struct {
			Temp_C float64 `json:"temp_c"`
		}{
			Temp_C: 20,
		},
	}
	if response2.Current.Temp_C != 20 {
		t.Errorf("Expected Temp_C to be 20, got '%f'", response2.Current.Temp_C)
	}

}
