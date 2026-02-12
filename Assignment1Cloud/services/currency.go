package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

type currencyResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// Function to collect the API Status
func getCurrencyStatus(apiURL string) int {
	resp, err := client.Get(apiURL)
	if err != nil {
		return 0 // cant get into service
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func GetExchangeStatus(baseCurrency string) (map[string]float64, error) {
	url := fmt.Sprintf("%s%s", CURRENCY_API, baseCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get currency API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get currency API: invalid status code: %d", resp.StatusCode)
	}

	var ratesData currencyResponse

	if err := json.NewDecoder(resp.Body).Decode(&ratesData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return ratesData.Rates, nil
}
