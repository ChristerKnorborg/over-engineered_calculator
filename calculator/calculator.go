package calculator

import (
	"errors"
	"math"
	"time"
)

type Calculator struct {
	History []HistoryEntry // History of operations performed
}

type HistoryEntry struct {
	Operand1  float64   // Left number in expression
	Operand2  float64   // Right number in expression
	Operation string    // +, -, *, /, %, ^
	Result    float64   // Result after applying the operation
	Timestamp time.Time // When the operation was performed
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

func (c *Calculator) saveToHistory(operation string, operand1, operand2, result float64) {
	entry := HistoryEntry{
		Operation: operation,
		Operand1:  operand1,
		Operand2:  operand2,
		Result:    result,
		Timestamp: time.Now(),
	}
	c.History = append(c.History, entry)
}

func (c *Calculator) GetHistory() []HistoryEntry {
	return c.History
}
