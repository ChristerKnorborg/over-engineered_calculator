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

type Storage interface {
	Save(entry HistoryEntry) error
	GetHistory() ([]HistoryEntry, error)
}

// Used to store history in memory for unit tests
type LocalStorage struct {
	History []HistoryEntry
}

// FirestoreStorage is used to store history in a Firestore database.
type FirestoreStorage struct {
	Client  *firestore.Client
	Context context.Context
}

// Save the history entry to the LocalStorage
func (storage *LocalStorage) Save(entry HistoryEntry) error {
	storage.History = append(storage.History, entry)
	return nil
}

// Get the history from the LocalStorage
func (storage *LocalStorage) GetHistory() ([]HistoryEntry, error) {
	if len(storage.History) == 0 {
		return nil, errors.New("no history found")
	}
	return storage.History, nil
}

// Save the history entry to Firestore database
func (storage *FirestoreStorage) Save(entry HistoryEntry) error {
	_, _, err := storage.Client.Collection("calculations").Add(storage.Context, map[string]interface{}{
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

	// Query the Firestore database by "timestamp". Probably add pagination given a real application
	iter := storage.Client.Collection("calculations").OrderBy("timestamp", firestore.Desc).Documents(storage.Context)

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
