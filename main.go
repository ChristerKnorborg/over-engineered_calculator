package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"overengineered_calculator/api"
	"overengineered_calculator/calculator"

	"cloud.google.com/go/firestore"      // Import the Firestore client
	firebase "firebase.google.com/go/v4" // Import the Firestore package
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

func initFirestore() error {
	var app *firebase.App
	var err error

	ctx := context.Background()

	// Check if running in a local environment with the Firestore emulator
	if emulatorHost := os.Getenv("FIRESTORE_EMULATOR_HOST"); emulatorHost != "" {
		// Connecting to Firestore emulator
		fmt.Println("Connecting to Firestore emulator at", emulatorHost)
		app, err = firebase.NewApp(ctx, &firebase.Config{
			ProjectID: "overengineered-calculato-2f35d",
		})
	} else {
		// Connect to the real Firestore service
		fmt.Println("Connecting to Firestore service...")
		opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
		app, err = firebase.NewApp(ctx, nil, opt)
	}

	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error initializing Firestore: %v", err)
	}

	return nil
}

func enableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(writer, request)
	})
}

func main() {

	ctx := context.Background()

	// Initialize Firestore
	if err := initFirestore(); err != nil {
		log.Fatalf("Firestore initialization failed: %v", err)
	}
	defer firestoreClient.Close()

	// Init calculator instance with Firestore storage
	firestoreStorage := &calculator.FirestoreStorage{
		Client:  firestoreClient,
		Context: ctx,
	}
	calc := calculator.Calculator{Storage: firestoreStorage}

	// Set the calculator instance in the API package
	api.SetCalculator(calc)

	// Create HTTP request multiplexer
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)
	handlerWithCors := enableCors(mux)

	// Start HTTP server on port 8080 for manual testing
	fmt.Println("Starting server on :8080...")
	fmt.Println("Firestore emulator:", os.Getenv("FIRESTORE_EMULATOR_HOST"))
	err := http.ListenAndServe(":8080", handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}

}
