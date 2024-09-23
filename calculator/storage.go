package calculator

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
)

type HistoryEntry struct {
	Operand1  float64   // Left operand in expression
	Operand2  float64   // Right operand in expression
	Operation string    // +, -, *, /, %, ^
	Result    float64   // Result after applying the operation
	Timestamp time.Time // When the operation was performed
}

type storage interface {
	save(entry HistoryEntry) error
	getHistory() ([]HistoryEntry, error)
	reset() error
}

// Used to store history in memory for unit tests
type LocalStorage struct {
	history []HistoryEntry
}

// FirestoreStorage is used to store history in a Firestore database.
type FirestoreStorage struct {
	client *firestore.Client
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		history: []HistoryEntry{},
	}
}

func NewFirestoreStorage(client *firestore.Client) *FirestoreStorage {
	return &FirestoreStorage{
		client: client,
	}
}

// Save the history entry to the LocalStorage
func (storage *LocalStorage) save(entry HistoryEntry) error {
	storage.history = append(storage.history, entry)
	return nil
}

// Get the history from the LocalStorage
func (storage *LocalStorage) getHistory() ([]HistoryEntry, error) {
	if len(storage.history) == 0 {
		return nil, errors.New("no history found")
	}
	return storage.history, nil
}

// Reset the history in the LocalStorage
func (storage *LocalStorage) reset() error {
	storage.history = []HistoryEntry{}
	return nil
}

// Save the history entry to Firestore database
func (storage *FirestoreStorage) save(entry HistoryEntry) error {

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
func (storage *FirestoreStorage) getHistory() ([]HistoryEntry, error) {

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
func (storage *FirestoreStorage) reset() error {

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
