// Package calculator provides the arithmetic operations (+, -, *, /, %, ^) and history tracking.
package calculator

import (
	"errors"
	"math"
)

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add takes two float64 operands, returns their sum, and saves the operation to history.
func (calc *Calculator) Add(operand1 float64, operand2 float64) float64 {

	result := operand1 + operand2
	return result
}

// Subtract takes two float64 operands, returns their difference, and saves the operation to history.
func (calc *Calculator) Subtract(operand1 float64, operand2 float64) float64 {

	result := operand1 - operand2
	return result
}

// Multiply takes two float64 operands, returns their product, and saves the operation to history.
func (calc *Calculator) Multiply(operand1 float64, operand2 float64) float64 {

	result := operand1 * operand2
	return result
}

// Divide takes two float64 operands, returns their quotient, and saves the operation to history.
// If dividing by zero, the function returns an error.
func (calc *Calculator) Divide(operand1 float64, operand2 float64) (float64, error) {

	if operand2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	result := operand1 / operand2
	return result, nil
}

// Modulo takes two float64 operands, returns the remainder of the division, and saves the operation to history.
// If modulo by zero, the function returns an error.
func (calc *Calculator) Modulo(operand1 float64, operand2 float64) (float64, error) {

	if operand2 == 0 {
		return 0, errors.New("cannot modulo by zero")
	}

	result := math.Mod(operand1, operand2) // Standard "%" operator does not work with floats
	return result, nil
}

func (calc *Calculator) Power(operand1 float64, operand2 float64) float64 {

	result := math.Pow(operand1, operand2)
	return result
}
