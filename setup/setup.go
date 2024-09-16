package setup

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

// Initialize Firestore connection using service account key
func InitFirestoreProduction() error {

	ctx := context.Background()

	// Connect to the real Firestore service
	fmt.Println("Connecting to Firestore service...")
	opt := option.WithCredentialsFile("secrets/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error initializing Firestore: %v", err)
	}

	return nil
}

// Emulator database for testing purposes
func InitFirestoreEmulator() error {
	var app *firebase.App
	var err error

	ctx := context.Background()

	if emulatorHost := os.Getenv("FIRESTORE_EMULATOR_HOST"); emulatorHost != "" {
		// Connect to Firestore emulator
		fmt.Println("Connecting to Firestore emulator at", emulatorHost)
		app, err = firebase.NewApp(ctx, &firebase.Config{
			ProjectID: "overengineered-calculato-2f35d",
		})
	}

	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	// Initialize Firestore emulator
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error initializing Firestore: %v", err)
	}

	return nil
}

// CORS on top of HTTP is needed to tell the client (browser)
// what HTTP requests it is allowed to make. For simplicity, I allow all origins "*".
// This is not secure for production use, instead it should be changed to the frontend URL.
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
