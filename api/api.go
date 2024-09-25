package api

import (
	"log"
	"overengineered_calculator/calculator"
	"overengineered_calculator/storage"
	"time"
)

type API struct {
	calculator *calculator.Calculator
	storage    storage.Storage
}

func NewAPI(calc *calculator.Calculator, storage storage.Storage) *API {
	return &API{
		calculator: calc,
		storage:    storage,
	}
}

// saveToHistory saves the operation and its result to the history.
func (api *API) saveToHistory(operation string, operand1, operand2, result float64) {

	entry := storage.HistoryEntry{
		Operation: operation,
		Operand1:  operand1,
		Operand2:  operand2,
		Result:    result,
		Timestamp: time.Now(),
	}

	err := api.storage.SaveOperation(entry)
	if err != nil {
		log.Printf("Failed to save history: %v", err)
	}
}
