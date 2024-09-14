package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"overengineered_calculator/calculator"
	"strconv"
)

var calc = calculator.Calculator{}

// Helper function to parse operands from the request as float64
func parseOperands(request *http.Request) (float64, float64, error) {
	operand1, err1 := strconv.ParseFloat(request.URL.Query().Get("operand1"), 64)
	operand2, err2 := strconv.ParseFloat(request.URL.Query().Get("operand2"), 64)
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid operands")
	}

	return operand1, operand2, nil
}

// Handler for Add operation
func AddHandler(writer http.ResponseWriter, request *http.Request) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Add(operand1, operand2)
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Handler for Subtract operation
func SubtractHandler(writer http.ResponseWriter, request *http.Request) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Subtract(operand1, operand2)
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Handler for Multiply operation
func MultiplyHandler(writer http.ResponseWriter, request *http.Request) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Multiply(operand1, operand2)
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Handler for Divide operation
func DivideHandler(writer http.ResponseWriter, request *http.Request) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calc.Divide(operand1, operand2)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Handler for Modulo operation
func ModuloHandler(writer http.ResponseWriter, request *http.Request) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calc.Modulo(operand1, operand2)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Handler for Power operation
func PowerHandler(writer http.ResponseWriter, request *http.Request) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Power(operand1, operand2)
	json.NewEncoder(writer).Encode(map[string]float64{"result": result})
}

// Handler for retrieving history
func HistoryHandler(writer http.ResponseWriter, request *http.Request) {
	history := calc.GetHistory()
	json.NewEncoder(writer).Encode(history)
}
