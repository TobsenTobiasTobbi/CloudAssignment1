package handler

import (
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
}

// Decides where to send user from information
func handleStatusGetRequest(w http.ResponseWriter, r *http.Request) {
	// Getting status from both API's, might have to change where this is
	restCountriesStatus := getAPIStatus(REST_COUNTRIES_API)
	currencyAPIStatus := getAPIStatus(CURRENCY_API)

	// Calculating uptime
	uptime := int(time.Since(startTime).Seconds())

	// Response
	response := map[string]interface{}{
		"restcountriesapi": restCountriesStatus,
		"currenciesapi":    currencyAPIStatus,
		"version":          "v1",
		"uptime":           uptime,
	}
}

	// Sending out a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

func getAPIStatus(apiURL string) int {
	respon, erro := http.Get(apiURL)
	if erro != nil {
		return 0 // cant get into service
	}
	defer respon.Body.Close()

	return respon.StatusCode
}

