package calculator

import (
	"context"
	"errors"
	"log"
	"math"
	"time"

	"cloud.google.com/go/firestore"
)

type Calculator struct {
	Storage Storage
}

type HistoryEntry struct {
	Operand1  float64   // Left number in expression
	Operand2  float64   // Right number in expression
	Operation string    // +, -, *, /, %, ^
	Result    float64   // Result after applying the operation
	Timestamp time.Time // When the operation was performed
}

type Storage interface {
	Save(entry HistoryEntry) error
	GetHistory() ([]HistoryEntry, error)
}

type LocalStorage struct {
	History []HistoryEntry
}

type FirestoreStorage struct {
	Client  *firestore.Client
	Context context.Context
}

func (calc *Calculator) Add(Operand1 float64, Operand2 float64) float64 {
	result := Operand1 + Operand2
	calc.saveToHistory("Add", Operand1, Operand2, result)
	return result
}

func (calc *Calculator) Subtract(Operand1 float64, Operand2 float64) float64 {
	result := Operand1 - Operand2
	calc.saveToHistory("Subtract", Operand1, Operand2, result)
	return result
}

func (calc *Calculator) Multiply(Operand1 float64, Operand2 float64) float64 {
	result := Operand1 * Operand2
	calc.saveToHistory("Multiply", Operand1, Operand2, result)
	return result
}

func (calc *Calculator) Divide(Operand1 float64, Operand2 float64) (float64, error) {
	if Operand2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := Operand1 / Operand2
	calc.saveToHistory("Divide", Operand1, Operand2, result)

	return result, nil
}

func (calc *Calculator) Modulo(Operand1 float64, Operand2 float64) (float64, error) {
	if Operand2 == 0 {
		return 0, errors.New("cannot modulo by zero")
	}

	result := math.Mod(Operand1, Operand2) // Standard "%"" operator does not work with floats
	calc.saveToHistory("Modulo", Operand1, Operand2, result)

	return result, nil
}
func (calc *Calculator) Power(Operand1 float64, Operand2 float64) float64 {
	result := math.Pow(Operand1, Operand2)
	calc.saveToHistory("Power", Operand1, Operand2, result)
	return result
}

func (calc *Calculator) saveToHistory(operation string, operand1, operand2, result float64) {
	entry := HistoryEntry{
		Operation: operation,
		Operand1:  operand1,
		Operand2:  operand2,
		Result:    result,
		Timestamp: time.Now(),
	}

	err := calc.Storage.Save(entry)
	if err != nil {
		log.Printf("Failed to save history: %v", err)
	}
}

func (calc *Calculator) GetHistory() ([]HistoryEntry, error) {
	return calc.Storage.GetHistory()
}

func (storage *LocalStorage) Save(entry HistoryEntry) error {
	storage.History = append(storage.History, entry)
	return nil
}

func (storage *LocalStorage) GetHistory() ([]HistoryEntry, error) {
	if len(storage.History) == 0 {
		return nil, errors.New("no history found")
	}
	return storage.History, nil
}

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

func (storage *FirestoreStorage) GetHistory() ([]HistoryEntry, error) {

	var history []HistoryEntry

	// Query the Firestore database by "timestamp". Probably add pagination given a real application
	iter := storage.Client.Collection("calculations").OrderBy("timestamp", firestore.Desc).Documents(storage.Context)

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var entry HistoryEntry
		err = doc.DataTo(&entry)
		if err != nil {
			return nil, err
		}
		history = append(history, entry)
	}

	return history, nil
}
