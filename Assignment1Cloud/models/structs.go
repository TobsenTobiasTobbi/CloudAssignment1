package models

type StatusResponse struct {
	RestCountriesAPI int `json: "restcountriesapi"`
	CurrenciesAPI    int
	Version          string
	Uptime           int64
}
