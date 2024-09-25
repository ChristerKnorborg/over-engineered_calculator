package storage

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
)

// FirestoreStorage is used to store history in a Firestore database.
type FirestoreStorage struct {
	client *firestore.Client
}

func NewFirestoreStorage(client *firestore.Client) *FirestoreStorage {
	return &FirestoreStorage{
		client: client,
	}
}

// Save the history entry to Firestore database
func (storage *FirestoreStorage) SaveOperation(entry HistoryEntry) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, _, err := storage.client.Collection("calculations").Add(ctx, map[string]interface{}{
		"operand1":  entry.Operand1,
		"operand2":  entry.Operand2,
		"operation": entry.Operation,
		"result":    entry.Result,
		"timestamp": entry.Timestamp,
	})
	return err
}

// The function GetHistory retrieves the history of calculations from the Firestore database sorted by newest operations
// first. It returns a slice of HistoryEntry structs.
func (storage *FirestoreStorage) GetHistory() ([]HistoryEntry, error) {

	var history []HistoryEntry

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query the Firestore database by "timestamp". Probably add pagination given a real application
	iter := storage.client.Collection("calculations").OrderBy("timestamp", firestore.Desc).Documents(ctx)

	for {
		document, err := iter.Next()
		if err != nil {
			break
		}

		var entry HistoryEntry
		err = document.DataTo(&entry)
		if err != nil {
			return nil, err
		}
		history = append(history, entry)
	}

	return history, nil
}

// Reset the history in the Firestore database
func (storage *FirestoreStorage) ResetHistory() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := storage.client.Collection("calculations").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		_, err = doc.Ref.Delete(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// RegisterUser stores the username and hashed password in Firestore.
func (storage *FirestoreStorage) RegisterUser(username string, password string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	// Check if the user already exists
	docRef := storage.client.Collection("users").Doc(username)
	doc, err := docRef.Get(ctx)
	if err == nil && doc.Exists() {
		return errors.New("user already exists")
	}

	// Store the username and hash of the password
	_, err = docRef.Set(ctx, map[string]interface{}{
		"username": username,
		"password": hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

// AuthenticateUser retrieves the stored hashed password and checks it against the hash of the provided password.
func (storage *FirestoreStorage) AuthenticateUser(username string, password string) error {

	// Create context with a 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the user from Firestore
	docRef := storage.client.Collection("users").Doc(username)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return errors.New("user not found")
	}

	// Get the stored hash og the password
	var user struct {
		Password string `firestore:"password"`
	}

	// Convert the Firestore document to a struct
	err = doc.DataTo(&user)
	if err != nil {
		return err
	}

	// Check if the provided password matches the stored hash
	correctPassword := checkPasswordHash(password, user.Password)
	if !correctPassword {
		return errors.New("invalid credentials")
	}

	return nil
}

// HashPassword hashes a plaintext password using bcrypt with a cost of 14.
// The cost increases the work factor with 2^cost, making it slower to brute force.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a plaintext password with a hashed password using bcrypt.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
