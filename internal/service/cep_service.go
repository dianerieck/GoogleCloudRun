package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dianerieck/Multithreading-Golang/internal/dto"
)

func fetchFromBrasilAPI(cep string, ch chan<- *apiResult) {
	start := time.Now()
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	resp, err := http.Get(url)
	if err != nil {
		ch <- &apiResult{data: fmt.Sprintf("Erro BrasilAPI: %v", err), source: "BrasilAPI", latency: time.Since(start)}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- &apiResult{data: fmt.Sprintf("Erro BrasilAPI: %v", err), source: "BrasilAPI", latency: time.Since(start)}
		return
	}

	var data dto.BrasilAPIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- &apiResult{data: fmt.Sprintf("Erro BrasilAPI: %v", err), source: "BrasilAPI", latency: time.Since(start)}
		return
	}

	ch <- &apiResult{data: &data, source: "BrasilAPI", latency: time.Since(start)}
}

func fetchFromViaCEP(cep string, ch chan<- *apiResult) {
	start := time.Now()
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		ch <- &apiResult{data: fmt.Sprintf("Erro ViaCEP: %v", err), source: "ViaCEP", latency: time.Since(start)}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- &apiResult{data: fmt.Sprintf("Erro ViaCEP: %v", err), source: "ViaCEP", latency: time.Since(start)}
		return
	}

	var data dto.ViaCEPResponse
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- &apiResult{data: fmt.Sprintf("Erro ViaCEP: %v", err), source: "ViaCEP", latency: time.Since(start)}
		return
	}

	if data.Erro {
		ch <- &apiResult{data: fmt.Sprintf("CEP não encontrado em ViaCEP"), source: "ViaCEP", latency: time.Since(start)}
		return
	}

	ch <- &apiResult{data: &data, source: "ViaCEP", latency: time.Since(start)}
}

type apiResult struct {
	data    interface{}
	source  string
	latency time.Duration
}

func FetchCEPRace(cep string) (interface{}, string, time.Duration) {
	ch := make(chan *apiResult, 2) // Buffer para ambas as APIs

	go fetchFromBrasilAPI(cep, ch)
	go fetchFromViaCEP(cep, ch)

	select {
	case result := <-ch:
		return result.data, result.source, result.latency
	case <-time.After(time.Second * 1):
		return nil, "", 0
	}
}
