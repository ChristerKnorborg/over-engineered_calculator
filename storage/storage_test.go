package storage

import (
	"testing"
)

func setupLocalStorage() *localStorage {
	return NewLocalStorage()
}

func TestLocalStorageSave(t *testing.T) {
	storage := setupLocalStorage()

	entry := HistoryEntry{
		Operand1:  5,
		Operand2:  3,
		Operation: "Add",
		Result:    8,
	}

	err := storage.SaveOperation(entry)
	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if len(storage.history) != 1 {
		t.Errorf("Expected 1 history entry but got %d", len(storage.history))
	}

	if storage.history[0].Operand1 != 5 {
		t.Errorf("Expected 5 but got %f", storage.history[0].Operand1)
	}

	if storage.history[0].Operand2 != 3 {
		t.Errorf("Expected 3 but got %f", storage.history[0].Operand2)
	}

	if storage.history[0].Operation != "Add" {
		t.Errorf("Expected Add but got %s", storage.history[0].Operation)
	}

	if storage.history[0].Result != 8 {
		t.Errorf("Expected 8 but got %f", storage.history[0].Result)
	}

}

func TestLocalStorageGetHistory(t *testing.T) {
	storage := setupLocalStorage()

	entry := HistoryEntry{
		Operand1:  5,
		Operand2:  3,
		Operation: "Add",
		Result:    8,
	}
	storage.SaveOperation(entry)

	history, err := storage.GetHistory()
	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if len(history) != 1 {
		t.Errorf("Expected 1 history entry but got %d", len(history))
	}

	if history[0].Operand1 != 5 {
		t.Errorf("Expected 5 but got %f", history[0].Operand1)
	}

	if history[0].Operand2 != 3 {

		t.Errorf("Expected 3 but got %f", history[0].Operand2)
	}

	if history[0].Operation != "Add" {
		t.Errorf("Expected Add but got %s", history[0].Operation)
	}

	if history[0].Result != 8 {
		t.Errorf("Expected 8 but got %f", history[0].Result)
	}
}

func TestLocalStorageResetHistory(t *testing.T) {
	storage := setupLocalStorage()

	entry := HistoryEntry{
		Operand1:  5,
		Operand2:  3,
		Operation: "Add",
		Result:    8,
	}
	storage.SaveOperation(entry)

	err := storage.ResetHistory()
	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if len(storage.history) != 0 {
		t.Errorf("Expected 0 history entry but got %d", len(storage.history))
	}
}

func TestLocalStorageResetHistoryManyEntries(t *testing.T) {
	storage := setupLocalStorage()

	for i := 0; i < 1000; i++ {

		entry := HistoryEntry{
			Operand1:  float64(i),
			Operand2:  float64(i + 5),
			Operation: "Add",
			Result:    8, // Just a placeholder value
		}
		storage.SaveOperation(entry)
	}

	err := storage.ResetHistory()
	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if len(storage.history) != 0 {
		t.Errorf("Expected 0 history entry but got %d", len(storage.history))
	}
}

func TestLocalStorageResetHistoryEmpty(t *testing.T) {
	storage := setupLocalStorage()

	err := storage.ResetHistory()
	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if len(storage.history) != 0 {
		t.Errorf("Expected 0 history entry but got %d", len(storage.history))
	}
}
