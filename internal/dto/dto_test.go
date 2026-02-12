package dto

import "testing"

func TestBrasilAPIResponse(t *testing.T) {
	response := &BrasilAPIResponse{
		Cep:          "01153000",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Centro",
		Street:       "Avenida Paulista",
		Service:      "brasilapi",
	}

	if response.Cep != "01153000" {
		t.Errorf("CEP esperado: 01153000, obtido: %s", response.Cep)
	}

	if response.City != "São Paulo" {
		t.Errorf("Cidade esperada: São Paulo, obtida: %s", response.City)
	}
}

func TestViaCEPResponse(t *testing.T) {
	response := &ViaCEPResponse{
		Cep:        "01153000",
		Logradouro: "Avenida Paulista",
		Bairro:     "Centro",
		Localidade: "São Paulo",
		Uf:         "SP",
		Ibge:       "3550308",
		Ddd:        "11",
		Erro:       false,
	}

	if response.Cep != "01153000" {
		t.Errorf("CEP esperado: 01153000, obtido: %s", response.Cep)
	}

	if response.Uf != "SP" {
		t.Errorf("UF esperado: SP, obtido: %s", response.Uf)
	}

	if response.Erro {
		t.Errorf("Erro não era esperado")
	}
}

// Removed TestResultStructure because generic Result was removed
