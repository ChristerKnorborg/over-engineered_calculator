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

// Initialize Firestore connection using service account key
func InitFirestore() (*firestore.Client, error) {

	ctx := context.Background()

	// Connect to the real Firestore service
	fmt.Println("Connecting to Firestore service...")
	opt := option.WithCredentialsFile("/app/secrets/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "overengineered-calculato-2f35d",
	}, opt)

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firestore: %v", err)
	}

	return firestoreClient, nil
}

// Emulator database for testing purposes
func InitFirestoreEmulator() (*firestore.Client, error) {
	var app *firebase.App
	var err error

	ctx := context.Background()

	// Connect to Firestore emulator if running
	if emulatorHost := os.Getenv("FIRESTORE_EMULATOR_HOST"); emulatorHost != "" {
		fmt.Println("Connecting to Firestore emulator at", emulatorHost)
		app, err = firebase.NewApp(ctx, &firebase.Config{
			ProjectID: "overengineered-calculato-2f35d",
		})
	}

	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	// Initialize Firestore emulator
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return firestoreClient, fmt.Errorf("error initializing Firestore: %v", err)
	}

	return firestoreClient, nil
}

// CORS on top of HTTP is needed to tell the client (browser)
// what HTTP requests it is allowed to make. For simplicity, I allow all origins "*".
// This is not secure for production use, instead it should be changed to the frontend URL.
// However, I will not change this as it might make complication for the postman tests.
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
