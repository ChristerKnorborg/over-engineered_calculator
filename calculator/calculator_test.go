package calculator

import (
	"math"
	"testing"
)

func setupCalculatorWithLocalStorage() *Calculator {
	calculator := NewCalculator()
	return calculator
}

func TestAddPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()
	result := calc.Add(10000, 10000)
	if result != 20000 {
		t.Errorf("Expected 20000 but got %f", result)
	}
}

func TestAddNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Add(-10000, -10000)
	if result != -20000 {
		t.Errorf("Expected -20000 but got %f", result)
	}
}

func TestAddPositiveAndNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Add(-10000, 10000)
	if result != 0 {
		t.Errorf("Expected 0 but got %f", result)
	}
}

func TestAddZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Add(10000, 0)
	if result != 10000 {
		t.Errorf("Expected 10000 but got %f", result)
	}
}

func TestAddToZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Add(0, 10000)
	if result != 10000 {
		t.Errorf("Expected 10000 but got %f", result)
	}
}

func TestSubtractPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Subtract(10000, 10000)
	if result != 0 {
		t.Errorf("Expected 0 but got %f", result)
	}
}

func TestSubtractNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Subtract(-10000, -10000)
	if result != 0 {
		t.Errorf("Expected 0 but got %f", result)
	}
}

func TestSubtractPositiveAndNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Subtract(-10000, 10000)
	if result != -20000 {
		t.Errorf("Expected -20000 but got %f", result)
	}
}

func TestSubtractNegativeAndPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Subtract(10000, -10000)
	if result != 20000 {
		t.Errorf("Expected 20000 but got %f", result)
	}
}

func TestSubtractZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Subtract(10000, 0)
	if result != 10000 {
		t.Errorf("Expected 10000 but got %f", result)
	}
}

func TestSubtractFromZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Subtract(0, 10000)
	if result != -10000 {
		t.Errorf("Expected -10000 but got %f", result)
	}
}

func TestMultiplyPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Multiply(10000, 10000)
	if result != 100000000 {
		t.Errorf("Expected 100000000 but got %f", result)
	}
}

func TestMultiplyNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Multiply(-10000, -10000)
	if result != 100000000 {
		t.Errorf("Expected 100000000 but got %f", result)
	}
}

func TestMultiplyPositiveAndNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Multiply(-10000, 10000)
	if result != -100000000 {
		t.Errorf("Expected -100000000 but got %f", result)
	}
}

func TestMultiplyNegativeAndPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Multiply(10000, -10000)
	if result != -100000000 {
		t.Errorf("Expected -100000000 but got %f", result)
	}
}

func TestMultiplyWithZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Multiply(10000, 0)
	if result != 0 {
		t.Errorf("Expected 0 but got %f", result)
	}
}

func TestMultiplyNegativeOperandWithZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Multiply(-10000, 0)
	if result != 0 {
		t.Errorf("Expected 0 but got %f", result)
	}
}

func TestDividePositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Divide(10000, 10000)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != 1 {
		t.Errorf("Expected 1 but got %f", result)
	}
}

func TestDivideNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Divide(-10000, -10000)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != 1 {
		t.Errorf("Expected 1 but got %f", result)
	}
}

func TestDividePositiveAndNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Divide(-10000, 10000)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != -1 {
		t.Errorf("Expected -1 but got %f", result)
	}
}

func TestDivideNegativeAndPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Divide(-10000, 10000)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != -1 {
		t.Errorf("Expected -1 but got %f", result)
	}
}

func TestDivideByZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	_, err := calc.Divide(10000, 0)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestModuloPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Modulo(87, 5)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != 2 {
		t.Errorf("Expected 2 but got %f", result)
	}
}

func TestModuloNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Modulo(-87, -5)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != -2 {
		t.Errorf("Expected -2 but got %f", result)
	}
}

func TestModuloPositiveAndNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Modulo(-87, 5)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != -2 {
		t.Errorf("Expected -2 but got %f", result)
	}
}

func TestModuloNegativeAndPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Modulo(87, -5)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	if result != 2 {
		t.Errorf("Expected 2 but got %f", result)
	}
}

func TestModuloWithFloats(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result, err := calc.Modulo(87.3, 5.5)

	if err != nil {
		t.Errorf("Expected nil but got %s", err)
	}

	// Adjust with a small epsilon to account for floating point imprecision
	expected := 4.8
	epsilon := 0.00001

	if math.Abs(result-expected) > epsilon {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestModuloByZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	_, err := calc.Modulo(10000, 0)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestPowerPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Power(5, 3)
	if result != 125 {
		t.Errorf("Expected 8 but got %f", result)
	}
}

func TestPowerNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Power(-5, -2) // -5^-2 = -1/25
	if result != 0.04 {
		t.Errorf("Expected 0.04 but got %f", result)
	}
}

func TestPowerPositiveAndNegativeOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Power(-5, 3)
	if result != -125.0 {
		t.Errorf("Expected -125 but got %f", result)
	}
}

func TestPowerNegativeAndPositiveOperands(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Power(5, -3) // 5^-3 = 1/125
	if result != 0.008 {
		t.Errorf("Expected 0.008 but got %f", result)
	}
}

func TestPowerWithZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Power(5, 0)
	if result != 1 {
		t.Errorf("Expected 1 but got %f", result)
	}
}

func TestPowerZeroToZero(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	result := calc.Power(0, 0)
	if result != 1 {
		t.Errorf("Expected 1 for 0^0 but got %f", result)
	}
}
