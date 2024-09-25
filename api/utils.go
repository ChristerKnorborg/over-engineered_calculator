package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Type definitions for passing functions as arguments to handlers more cleanly
type calculatorOperation func(a, b float64) float64
type calculatorOperationWithError func(a, b float64) (float64, error)

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
