package main

import (
	"net/http"
	"log"
	"time"
	"os"

	"Assignment1Cloud/handlers"
)

func main () {

	// Handle port assignment (either based on environment variable, )
	port := Getenv("PORT")
	if port == "" {
		log.Println("$PORT har not been set. Default: 8080")
		port = "8080"
	}

	// Instantiate new router
	router := http.NewServeMux()

	// Set up and attach handler endpoints to router


	// Start server
	
}