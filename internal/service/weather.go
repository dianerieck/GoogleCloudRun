package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dianerieck/GoogleCloudRun/internal/dto"
)

func GetTemperatura(cidade string, apiKey string) (float64, error) {
	if apiKey == "" {
		return 0, fmt.Errorf("missing API key")
	}

	apiKey = strings.TrimSpace(apiKey)
	if strings.HasPrefix(strings.ToLower(apiKey), "bearer ") {
		apiKey = strings.TrimSpace(apiKey[len("bearer "):])
	}

	client := &http.Client{Timeout: 5 * time.Second}
	q := url.QueryEscape(cidade)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, q)

	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("weather API error: %s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data dto.WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data.Current.Temp_C, nil

}
