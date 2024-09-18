package calculator

import (
	"math"
	"testing"
)

func setupCalculatorWithLocalStorage() *Calculator {

	storage := NewLocalStorage()
	calculator := NewCalculator(storage)
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

func TestLocalStorageSave(t *testing.T) {
	storage := &LocalStorage{}
	entry := HistoryEntry{
		Operand1:  5,
		Operand2:  3,
		Operation: "Add",
		Result:    8,
	}

	err := storage.save(entry)
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
	storage := &LocalStorage{}
	entry := HistoryEntry{
		Operand1:  5,
		Operand2:  3,
		Operation: "Add",
		Result:    8,
	}
	storage.save(entry)

	history, err := storage.getHistory()
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

func TestCalculatorWithHistorySingleOperation(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	calc.Subtract(5, 3)

	history, _ := calc.GetHistory()
	if len(history) != 1 {
		t.Errorf("Expected 2 history entries but got %d", len(history))
	}
	// Check operands
	if history[0].Operand1 != 5 || history[0].Operand2 != 3 {
		t.Errorf("Expected operands 5 and 3 but got %f and %f", history[0].Operand1, history[0].Operand2)
	}
	// Check operation and result
	if history[0].Operation != "Subtract" || history[0].Result != 2 {
		t.Errorf("Expected Subtract operation with result 2 but got %s with result %f", history[0].Operation, history[0].Result)
	}
}

func TestCalculatorWithHistoryAllOperations(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	calc.Add(1, 2)
	calc.Subtract(5, 3)
	calc.Multiply(3, 4)
	calc.Divide(10, 2)
	calc.Modulo(10, 3)
	calc.Power(2, 3)

	history, _ := calc.GetHistory()
	if len(history) != 6 {
		t.Errorf("Expected 6 history entries but got %d", len(history))
	}
	if history[0].Operation != "Add" || history[0].Result != 3 {
		t.Errorf("Expected Add operation with result 3 but got %s with result %f", history[0].Operation, history[0].Result)
	}
	if history[1].Operation != "Subtract" || history[1].Result != 2 {
		t.Errorf("Expected Subtract operation with result 2 but got %s with result %f", history[1].Operation, history[1].Result)
	}
	if history[2].Operation != "Multiply" || history[2].Result != 12 {
		t.Errorf("Expected Multiply operation with result 12 but got %s with result %f", history[2].Operation, history[2].Result)
	}
	if history[3].Operation != "Divide" || history[3].Result != 5 {
		t.Errorf("Expected Divide operation with result 5 but got %s with result %f", history[3].Operation, history[3].Result)
	}
	if history[4].Operation != "Modulo" || history[4].Result != 1 {
		t.Errorf("Expected Modulo operation with result 1 but got %s with result %f", history[4].Operation, history[4].Result)
	}
	if history[5].Operation != "Power" || history[5].Result != 8 {
		t.Errorf("Expected Power operation with result 8 but got %s with result %f", history[5].Operation, history[5].Result)
	}
}

func TestEmptyHistory(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	history, _ := calc.GetHistory()

	if len(history) != 0 {
		t.Errorf("Expected empty history but got %d entries", len(history))
	}
}

func TestLargeHistory(t *testing.T) {

	calc := setupCalculatorWithLocalStorage()

	// Add 10000 entries to the history with Add operation
	for i := 0; i < 10000; i++ {
		calc.Add(float64(i), float64(i+1))
	}

	// Check all entries are present
	history, _ := calc.GetHistory()
	if len(history) != 10000 {
		t.Errorf("Expected 150 history entries but got %d", len(history))
	}

	// Check random elements for correctness
	if history[55].Operand1 != 55 || history[55].Operand2 != 56 || history[55].Result != 111 {
		t.Errorf("First entry in history is incorrect. Got operands: %f, %f and result %f", history[55].Operand1, history[55].Operand2, history[111].Result)
	}

	if history[149].Operand1 != 149 || history[149].Operand2 != 150 || history[149].Result != 299 {
		t.Errorf("Last entry in history is incorrect. Got operands: %f, %f and result %f", history[149].Operand1, history[149].Operand2, history[149].Result)
	}

	if history[9999].Operand1 != 9999 || history[9999].Operand2 != 10000 || history[9999].Result != 19999 {
		t.Errorf("Last entry in history is incorrect. Got operands: %f, %f and result %f", history[9999].Operand1, history[9999].Operand2, history[9999].Result)
	}
}
