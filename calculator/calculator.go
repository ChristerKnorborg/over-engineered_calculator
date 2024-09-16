// Package calculator provides the arithmetic operations (+, -, *, /, %, ^) and history tracking.
package calculator

import (
	"errors"
	"log"
	"math"
	"time"
)

type Calculator struct {
	Storage Storage
}

// Add takes two float64 operands, returns their sum, and saves the operation to history.
func (calc *Calculator) Add(Operand1 float64, Operand2 float64) float64 {
	result := Operand1 + Operand2
	calc.saveToHistory("Add", Operand1, Operand2, result)
	return result
}

// Subtract takes two float64 operands, returns their difference, and saves the operation to history.
func (calc *Calculator) Subtract(Operand1 float64, Operand2 float64) float64 {
	result := Operand1 - Operand2
	calc.saveToHistory("Subtract", Operand1, Operand2, result)
	return result
}

// Multiply takes two float64 operands, returns their product, and saves the operation to history.
func (calc *Calculator) Multiply(Operand1 float64, Operand2 float64) float64 {
	result := Operand1 * Operand2
	calc.saveToHistory("Multiply", Operand1, Operand2, result)
	return result
}

// Divide takes two float64 operands, returns their quotient, and saves the operation to history.
// If dividing by zero, the function returns an error.
func (calc *Calculator) Divide(Operand1 float64, Operand2 float64) (float64, error) {
	if Operand2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := Operand1 / Operand2
	calc.saveToHistory("Divide", Operand1, Operand2, result)

	return result, nil
}

// Modulo takes two float64 operands, returns the remainder of the division, and saves the operation to history.
// If modulo by zero, the function returns an error.
func (calc *Calculator) Modulo(Operand1 float64, Operand2 float64) (float64, error) {
	if Operand2 == 0 {
		return 0, errors.New("cannot modulo by zero")
	}

	result := math.Mod(Operand1, Operand2) // Standard "%" operator does not work with floats
	calc.saveToHistory("Modulo", Operand1, Operand2, result)

	return result, nil
}
func (calc *Calculator) Power(Operand1 float64, Operand2 float64) float64 {
	result := math.Pow(Operand1, Operand2)
	calc.saveToHistory("Power", Operand1, Operand2, result)
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

	err := calc.Storage.Save(entry)
	if err != nil {
		log.Printf("Failed to save history: %v", err)
	}
}

// GetHistory retrieves the history of calculations from the storage.
func (calc *Calculator) GetHistory() ([]HistoryEntry, error) {
	return calc.Storage.GetHistory()
}
