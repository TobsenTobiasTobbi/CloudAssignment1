package models

type CountryInfo struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population string            `json:"population"`
	Area       string            `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
}
