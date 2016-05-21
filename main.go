package main

import (
	"log"
	"net/http"
)

func main() {
	// Setup the router
	router := NewRouter()

	// Listen on port :8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
