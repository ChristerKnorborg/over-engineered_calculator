// The package api contains the handlers for the calculator API i.e. th

package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"overengineered_calculator/calculator"
	"strconv"
)

type API struct {
	calculator *calculator.Calculator
}

func NewAPI(calculator *calculator.Calculator) *API {
	return &API{
		calculator: calculator,
	}
}

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

// Generic handler for operations that return no error (Add, Subtract, Multiply, Power)
func (api *API) operationHandler(writer http.ResponseWriter, request *http.Request, operation calculatorOperation) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := operation(operand1, operand2)
	writeResultJSON(writer, result)
}

// Generic handler for operations that return an error (divide, modulo)
func (api *API) operationHandlerWithError(writer http.ResponseWriter, request *http.Request, operation calculatorOperationWithError) {
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
func (api *API) addHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, api.calculator.Add)
}

// Handler for Subtract operation
func (api *API) subtractHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, api.calculator.Subtract)
}

// Handler for Multiply operation
func (api *API) multiplyHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, api.calculator.Multiply)
}

// Handler for Divide operation
func (api *API) divideHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandlerWithError(writer, request, api.calculator.Divide)
}

// Handler for Modulo operation
func (api *API) moduloHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandlerWithError(writer, request, api.calculator.Modulo)
}

// Handler for Power operation
func (api *API) powerHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, api.calculator.Power)
}

// Handler for retrieving history
func (api *API) historyHandler(writer http.ResponseWriter, request *http.Request) {
	history, err := api.calculator.GetHistory()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(history)
}
