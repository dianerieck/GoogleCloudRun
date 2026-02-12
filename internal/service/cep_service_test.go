package service

import (
	"testing"
	"time"

	"github.com/dianerieck/CURSO-GO/11-DESAFIO-CEP/internal/dto"
)

func TestFetchCEPRaceWithValidCEP(t *testing.T) {
	data, source, latency := FetchCEPRace("01153000")

	if data == nil {
		t.Error("Resultado não deveria ser nil para CEP válido")
	}

	if source == "" {
		t.Error("Source não deveria estar vazio")
	}

	if source != "BrasilAPI" && source != "ViaCEP" {
		t.Errorf("Source deve ser BrasilAPI ou ViaCEP, obtido: %s", source)
	}

	if latency <= 0 {
		t.Error("Latency deve ser maior que zero")
	}
}

func TestFetchCEPRaceWithInvalidCEP(t *testing.T) {
	data, _, latency := FetchCEPRace("00000000")

	// A resposta pode ser um erro string ou um resultado válido
	if data == nil {
		t.Error("Resultado não deveria ser nil")
	}

	// Verifica se a latência foi registrada
	if latency <= 0 {
		t.Error("Latency deve ser maior que zero")
	}
}

// TestFetchCEPRaceResponseTime testa o timeout de 1 segundo
// Nota: Este teste pode falhar em conexões muito lentas
func TestFetchCEPRaceResponseTime(t *testing.T) {
	data, _, latency := FetchCEPRace("01153000")

	if data != nil {
		// Verifica se a resposta veio em menos de 1 segundo
		if latency >= time.Second {
			t.Errorf("Latency deve ser menor que 1 segundo, obtido: %v", latency)
		}
	}
}

// TestFetchCEPRaceDataTypes testa se os tipos de dados retornados são corretos
func TestFetchCEPRaceDataTypes(t *testing.T) {
	data, _, _ := FetchCEPRace("01153000")

	if data == nil {
		t.Fatal("Resultado não deveria ser nil")
	}

	switch data := data.(type) {
	case *dto.BrasilAPIResponse:
		if data == nil {
			t.Error("BrasilAPIResponse não deveria ser nil")
		}
		if data.Cep == "" {
			t.Error("CEP não deveria estar vazio em BrasilAPIResponse")
		}
	case *dto.ViaCEPResponse:
		if data == nil {
			t.Error("ViaCEPResponse não deveria ser nil")
		}
		if data.Cep == "" {
			t.Error("CEP não deveria estar vazio em ViaCEPResponse")
		}
	case string:
		// Pode ser uma mensagem de erro
		t.Logf("Resposta foi erro: %s", data)
	default:
		t.Errorf("Tipo de dados inesperado: %T", data)
	}
}
