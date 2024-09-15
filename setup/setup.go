package setup

import (
	"context"
	"fmt"
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
		// Connecting to Firestore emulator
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
