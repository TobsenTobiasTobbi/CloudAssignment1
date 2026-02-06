package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func StatusHandler (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	


}