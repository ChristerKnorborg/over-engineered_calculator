package main

import (
	"fmt"
	"log"
	"net/http"
	"overengineered_calculator/api"
)

func main() {

	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	// Start HTTP server on port 8080 for manual testing
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
