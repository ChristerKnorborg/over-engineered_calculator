package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"overengineered_calculator/api"
	"overengineered_calculator/calculator"
	"overengineered_calculator/setup"
)

func main() {

	ctx := context.Background()

	// Initialize Firestore
	err := setup.InitFirestoreEmulator()
	if err != nil {
		log.Fatalf("Firestore initialization failed: %v", err)
	}
	defer setup.FirestoreClient.Close()

	// Init calculator instance with Firestore storage
	firestoreStorage := &calculator.FirestoreStorage{
		Client:  setup.FirestoreClient,
		Context: ctx,
	}
	calc := calculator.Calculator{Storage: firestoreStorage}

	// Set the calculator instance in the API package
	api.SetCalculator(calc)

	// Create HTTP request multiplexer
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)
	handlerWithCors := setup.EnableCORS(mux)

	// Start HTTP server on port 8080 for manual testing
	fmt.Println("Starting server on :8080...")
	err = http.ListenAndServe(":8080", handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}
