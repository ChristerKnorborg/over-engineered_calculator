// The package api contains the handlers for the calculator API i.e. th

package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"overengineered_calculator/calculator"
	"strconv"
)

var calc = calculator.Calculator{}

type calculatorOperation func(a, b float64) float64
type calculatorOperationWithError func(a, b float64) (float64, error)

// Function to set the calculator instance to keep compatability
// with both Local Function calls and Local Storage for unit tests
// and Firestore for production
func SetCalculator(calculator calculator.Calculator) {
	calc = calculator
}

// Helper function to parse operands from the request as float64
func parseOperands(request *http.Request) (float64, float64, error) {
	operand1, err1 := strconv.ParseFloat(request.URL.Query().Get("operand1"), 64)
	operand2, err2 := strconv.ParseFloat(request.URL.Query().Get("operand2"), 64)
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid operands")
	}

	return operand1, operand2, nil
}

// Helper function to write JSON response with the result
func writeResultJSON(writer http.ResponseWriter, result float64) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Generic handler for operations that return no error (Add, Subtract, Multiply, Power)
func operationHandler(writer http.ResponseWriter, request *http.Request, operation calculatorOperation) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := operation(operand1, operand2)
	writeResultJSON(writer, result)
}

// Generic handler for operations that return an error (divide, modulo)
func operationHandlerWithError(writer http.ResponseWriter, request *http.Request, operation calculatorOperationWithError) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := operation(operand1, operand2)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writeResultJSON(writer, result)
}

// Handler for Add operation
func AddHandler(writer http.ResponseWriter, request *http.Request) {
	operationHandler(writer, request, calc.Add)
}

// Handler for Subtract operation
func SubtractHandler(writer http.ResponseWriter, request *http.Request) {
	operationHandler(writer, request, calc.Subtract)
}

// Handler for Multiply operation
func MultiplyHandler(writer http.ResponseWriter, request *http.Request) {
	operationHandler(writer, request, calc.Multiply)
}

// Handler for Divide operation
func DivideHandler(writer http.ResponseWriter, request *http.Request) {
	operationHandlerWithError(writer, request, calc.Divide)
}

// Handler for Modulo operation
func ModuloHandler(writer http.ResponseWriter, request *http.Request) {
	operationHandlerWithError(writer, request, calc.Modulo)
}

// Handler for Power operation
func PowerHandler(writer http.ResponseWriter, request *http.Request) {
	operationHandler(writer, request, calc.Power)
}

// Handler for retrieving history
func HistoryHandler(writer http.ResponseWriter, request *http.Request) {
	history, err := calc.Storage.GetHistory()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(history)
}
