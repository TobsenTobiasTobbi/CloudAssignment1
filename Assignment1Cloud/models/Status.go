package models

type StatusResponse struct {
	RestCountriesAPI int    `json:"restcountriesapi"`
	CurrenciesAPI    int    `json:"currenciesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}
