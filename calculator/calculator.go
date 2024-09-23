// Package calculator provides the arithmetic operations (+, -, *, /, %, ^) and history tracking.
package calculator

import (
	"errors"
	"log"
	"math"
	"sync"
	"time"
)

type Calculator struct {
	storage storage
	mutex   sync.Mutex
}

func NewCalculator(storage storage) *Calculator {
	return &Calculator{
		storage: storage,
	}
}

// Add takes two float64 operands, returns their sum, and saves the operation to history.
func (calc *Calculator) Add(operand1 float64, operand2 float64) float64 {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	result := operand1 + operand2
	calc.saveToHistory("Add", operand1, operand2, result)
	return result
}

// Subtract takes two float64 operands, returns their difference, and saves the operation to history.
func (calc *Calculator) Subtract(operand1 float64, operand2 float64) float64 {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	result := operand1 - operand2
	calc.saveToHistory("Subtract", operand1, operand2, result)
	return result
}

// Multiply takes two float64 operands, returns their product, and saves the operation to history.
func (calc *Calculator) Multiply(operand1 float64, operand2 float64) float64 {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	result := operand1 * operand2
	calc.saveToHistory("Multiply", operand1, operand2, result)
	return result
}

// Divide takes two float64 operands, returns their quotient, and saves the operation to history.
// If dividing by zero, the function returns an error.
func (calc *Calculator) Divide(operand1 float64, operand2 float64) (float64, error) {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	if operand2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := operand1 / operand2
	calc.saveToHistory("Divide", operand1, operand2, result)

	return result, nil
}

// Modulo takes two float64 operands, returns the remainder of the division, and saves the operation to history.
// If modulo by zero, the function returns an error.
func (calc *Calculator) Modulo(operand1 float64, operand2 float64) (float64, error) {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	if operand2 == 0 {
		return 0, errors.New("cannot modulo by zero")
	}

	result := math.Mod(operand1, operand2) // Standard "%" operator does not work with floats
	calc.saveToHistory("Modulo", operand1, operand2, result)

	return result, nil
}
func (calc *Calculator) Power(operand1 float64, operand2 float64) float64 {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	result := math.Pow(operand1, operand2)
	calc.saveToHistory("Power", operand1, operand2, result)
	return result
}

// saveToHistory saves the operation and its result to the history.
func (calc *Calculator) saveToHistory(operation string, operand1, operand2, result float64) {

	entry := HistoryEntry{
		Operation: operation,
		Operand1:  operand1,
		Operand2:  operand2,
		Result:    result,
		Timestamp: time.Now(),
	}

	err := calc.storage.save(entry)
	if err != nil {
		log.Printf("Failed to save history: %v", err)
	}
}

// GetHistory retrieves the history of calculations from the storage.
func (calc *Calculator) GetHistory() ([]HistoryEntry, error) {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	return calc.storage.getHistory()
}

// ResetHistory clears the history of calculations from the storage.
func (calc *Calculator) ResetHistory() error {

	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	return calc.storage.reset()
}
