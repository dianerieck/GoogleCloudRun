package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dianerieck/GoogleCloudRun/internal/dto"
)

func GetCidade(cep string) (string, error) {

	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data dto.ViaCEPResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("can not find zipcode")
	}

	return data.Localidade, nil
}
