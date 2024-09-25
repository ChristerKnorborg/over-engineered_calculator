package main

import (
	"fmt"
	"log"
	"net/http"
	"overengineered_calculator/api"
	"overengineered_calculator/calculator"
	"overengineered_calculator/setup"
	"overengineered_calculator/storage"
)

func main() {

	// Initialize Firestore
	firestoreClient, err := setup.InitFirestoreEmulator()
	if err != nil {
		log.Fatalf("Firestore initialization failed: %v", err)
	}
	defer firestoreClient.Close()

	// Initialize Calculator with Firestore storage for API
	firestoreStorage := storage.NewFirestoreStorage(firestoreClient)
	calc := calculator.NewCalculator()
	api := api.NewAPI(calc, firestoreStorage)

	// Create HTTP request multiplexer
	multiplexer := http.NewServeMux()
	api.RegisterRoutes(multiplexer)
	handlerWithCors := setup.EnableCORS(multiplexer)

	fmt.Println("Starting server on :8080...")
	err = http.ListenAndServe(":8080", handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}
