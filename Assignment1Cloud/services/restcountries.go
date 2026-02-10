package services

import (
	"Assigment1Cloud/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCountryByCode Gets country information by the 2-letter country code, and returns error
// or the map containing the relevant country data
func GetCountryByCode(countryCode string) (*models.CountryInfo, error) {
	url := fmt.Sprintf("%s/alpha/%s", REST_COUNTRIES_API, countryCode)

	// GET request to the REST Countries API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// Closes the function when response body exits
	defer resp.Body.Close()

	// Checks if the API returns a usable status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("REST Countries API returned status %d", resp.StatusCode)
	}

	// Making JSON response into a slice (to work with it in my code)
	var countries []models.CountryInfo
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Verifying that at least one country was returned
	if len(countries) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	// Returning the only country as response
	return &countries[0], nil
}

// GetRestCountriesStatus Checks if the REST Countries API is reachable
func GetRestCountriesStatus() int {
	// Attempting to get the REST Countries API base URL
	resp, err := http.Get(REST_COUNTRIES_API)
	if err != nil {
		// service unreachable
		return 0
	}
	// Ensure that the response body closes
	defer resp.Body.Close()

	// Returning the actual HTTP status code
	return resp.StatusCode
}
