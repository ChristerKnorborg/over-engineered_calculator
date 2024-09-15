package api_test

import (
	"net/http"
	"net/http/httptest"
	"overengineered_calculator/api"
	"strings"
	"testing"
)

// TestAddHandler checks the addition handler with the two operands:
// operand1 = 10, operand2 = 5. The response should be 15.
func TestAddHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/add?operand1=10&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.AddHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check header
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", responseRecorder.Header().Get("Content-Type"))
	}

	// Check the result directly from the response body
	expected := `{"result":15}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

// TestAddHandlerWithInvalidOperands checks how the addition handler handles invalid operands:
// operand1 = invalid, operand2 = 5. Response should be 400 status code.
func TestAddHandlerWithInvalidOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/add?operand1=invalid&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.AddHandler(responseRecorder, request)

	// Check for bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestAddHandlerWithMissingOperands checks the addition handler with missing operands.
// operand1 = 10, operand2 is missing. Response should be 400 status code.
func TestAddHandlerWithMissingOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/add?operand1=10", nil)
	responseRecorder := httptest.NewRecorder()

	api.AddHandler(responseRecorder, request)

	// Check for the bad request status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestSubtractHandler checks the subtraction handler with the two operands:
// operand1 = 10, operand2 = 5. The response should be 5.
func TestSubtractHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/subtract?operand1=10&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.SubtractHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check header
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", responseRecorder.Header().Get("Content-Type"))
	}

	// Check the result directly from the response body
	expected := `{"result":5}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

// TestSubtractHandlerWithInvalidOperands checks the subtraction handler with invalid operands:
// operand1 = 10, operand2 = invalid. Response should be 400 status code.
func TestSubtractHandlerWithInvalidOperands(t *testing.T) {

	request := httptest.NewRequest("GET", "/subtract?operand1=10&operand2=invalid", nil)
	responseRecorder := httptest.NewRecorder()

	api.SubtractHandler(responseRecorder, request)

	// Check bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestSubtractHandlerWithMissingOperands checks the subtract handler with missing operands.
// operand1 = 10, operand2 is missing. Response should be 400 status code.
func TestSubtractHandlerWithMissingOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/subtract?operand1=10", nil)
	responseRecorder := httptest.NewRecorder()

	api.SubtractHandler(responseRecorder, request)

	// Check for the bad request status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestMultiplyHandler checks the multiplication handler with the two operands:
// operand1 = 10, operand2 = 5. The response should be 50.
func TestMultiplyHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/multiply?operand1=10&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.MultiplyHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check header
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", responseRecorder.Header().Get("Content-Type"))
	}

	// Check the result directly from the response body
	expected := `{"result":50}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

// TestMultiplyHandlerWithInvalidOperands checks the multiply handler with invalid operands:
// operand1 = invalid, operand2 = 45. Response should be 400 status code.
func TestMultiplyHandlerWithInvalidOperands(t *testing.T) {

	request := httptest.NewRequest("GET", "/multiply?operand1=invalid&operand2=45", nil)
	responseRecorder := httptest.NewRecorder()

	api.MultiplyHandler(responseRecorder, request)

	// Check bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestMultiplyHandlerWithMissingOperands checks the multiply handler with missing operands.
// operand1 is missing, operand2 = 10. Response should be 400 status code.
func TestMultiplyHandlerWithMissingOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/multiply?operand2=10", nil)
	responseRecorder := httptest.NewRecorder()

	api.MultiplyHandler(responseRecorder, request)

	// Check for the bad request status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestDivideHandler checks the division handler with the two operands:
// operand1 = 10, operand2 = 5. The response should be 2.
func TestDivideHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/divide?operand1=10&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.DivideHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check header
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", responseRecorder.Header().Get("Content-Type"))
	}

	// Check the result directly from the response body
	expected := `{"result":2}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

// TestDivideHandlerWithInvalidOperands checks the divide handler with invalid operands:
// operand1 = invalid, operand2 = 45. Response should be 400 status code.
func TestDivideHandlerWithInvalidOperands(t *testing.T) {

	request := httptest.NewRequest("GET", "/subtract?operand1=invalid&operand2=45", nil)
	responseRecorder := httptest.NewRecorder()

	api.DivideHandler(responseRecorder, request)

	// Check bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestDivideHandlerWithMissingOperands checks the divide handler with missing operands.
// operand1 is missing, operand2 = 10. Response should be 400 status code.
func TestDivideHandlerWithMissingOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/divide?operand2=10", nil)
	responseRecorder := httptest.NewRecorder()

	api.DivideHandler(responseRecorder, request)

	// Check for the bad request status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestDivideHandlerWithInvalidOperands checks the divide handler with division by zero:
// operand1 = 45, operand2 = 0. Response should be 400 status code.
func TestDivideHandlerWithDivisionByZero(t *testing.T) {

	request := httptest.NewRequest("GET", "/multiply?operand1=10&operand2=0", nil)
	responseRecorder := httptest.NewRecorder()

	api.DivideHandler(responseRecorder, request)

	// Check status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "cannot divide by zero\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestModuloHandler checks the modulo handler with the two operands:
// operand1 = 11, operand2 = 5. The response should be 1.
func TestModuloHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/modulo?operand1=11&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.ModuloHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check header
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", responseRecorder.Header().Get("Content-Type"))
	}

	// Check the result directly from the response body
	expected := `{"result":1}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

// TestModuloHandlerWithInvalidOperands checks the modulo handler with invalid operands:
// operand1 = 45, operand2 = invalid. Response should be 400 status code.
func TestModuloHandlerWithInvalidOperands(t *testing.T) {

	request := httptest.NewRequest("GET", "/modulo?operand1=45&operand2=invalid", nil)
	responseRecorder := httptest.NewRecorder()

	api.ModuloHandler(responseRecorder, request)

	// Check bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestModuloHandlerWithMissingOperands checks the modulo handler with missing operands.
// operand1 = 45, operand2 is missing. Response should be 400 status code.
func TestModuloHandlerWithMissingOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/modulo?operand1=45", nil)
	responseRecorder := httptest.NewRecorder()

	api.ModuloHandler(responseRecorder, request)

	// Check for bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestModuloHandlerWithDivisionByZero checks the modulo handler with division by zero:
// operand1 = 45, operand2 = 0. Response should be 400 status code.
func TestModuloHandlerWithModuloByZero(t *testing.T) {

	request := httptest.NewRequest("GET", "/modulo?operand1=45&operand2=0", nil)
	responseRecorder := httptest.NewRecorder()

	api.ModuloHandler(responseRecorder, request)

	// Check status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "cannot modulo by zero\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}

// TestPowerHandler checks the power handler with the two operands:
// operand1 = 2, operand2 = 3. The response should be 8.
func TestPowerHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/power?operand1=2&operand2=3", nil)
	responseRecorder := httptest.NewRecorder()

	api.PowerHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check header
	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", responseRecorder.Header().Get("Content-Type"))
	}

	// Check the result directly from the response body
	expected := `{"result":8}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

// TestPowerHandlerWithInvalidOperands checks the power handler with invalid operands:
// operand1 = invalid, operand2 = 3. Response should be 400 status code.
func TestPowerHandlerWithInvalidOperands(t *testing.T) {
	request := httptest.NewRequest("GET", "/power?operand1=invalid&operand2=3", nil)
	responseRecorder := httptest.NewRecorder()

	api.PowerHandler(responseRecorder, request)

	// Check for bad status code
	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", responseRecorder.Code)
	}

	// Check error message is returned
	expectedError := "invalid operands\n"
	actualError := responseRecorder.Body.String()
	if actualError != expectedError {
		t.Fatalf("expected %v, got %v", expectedError, actualError)
	}
}
