package models

type ExchangeResponse struct {
	Country       string               `json:"country"`
	Basecurrency  string               `json:"basecurrency"`
	ExchangeRates []map[string]float64 `json:"exchange_rates"`
}
