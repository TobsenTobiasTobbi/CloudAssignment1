package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func InfoHandler (w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST Method '" + r.Method + "' not supported. Currently only '" + http.MethodGet +
					"' and '" + http.MethodPost + "' are supported.", http.StatusNotImplemented)
	}
	
}