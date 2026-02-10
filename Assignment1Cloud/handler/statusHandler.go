package handler

import (
	"Assigment1Cloud/models"
	"encoding/json"
	"net/http"
	"time"
)

// Cotrolling the time when the service starts
var startTime time.Time

func init() {
	startTime = time.Now()
}

// Status handler, either let through or not
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handleStatusGetRequest(w, r)
}

// Decides where to send user from information
func handleStatusGetRequest(w http.ResponseWriter, r *http.Request) {
	// Getting status from both API's, might have to change where this is
	restCountriesStatus := getAPIStatus(REST_COUNTRIES_API)
	currencyAPIStatus := getAPIStatus(CURRENCY_API)

	// Calculating uptime
	uptime := int(time.Since(startTime).Seconds())

	// Response, gets it from models StatusResponse
	response := models.StatusResponse{
		RestCountriesAPI: restCountriesStatus,
		CurrenciesAPI:    currencyAPIStatus,
		Version:          "v1",
		Uptime:           int64(uptime),
	}

	// Sending out a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)

}

// Function to collect the API Status
func getAPIStatus(apiURL string) int {
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0 // cant get into service
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
