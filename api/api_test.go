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

	// "GET", "/add?operand1="10"&operand2="5", nil
	request := httptest.NewRequest("GET", "/add?operand1=10&operand2=5", nil)
	responseRecorder := httptest.NewRecorder()

	api.AddHandler(responseRecorder, request)

	// Check the status code
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", responseRecorder.Code)
	}

	// Check the result directly from the response body
	expected := `{"result":15}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
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

	// Check the result directly from the response body
	expected := `{"result":5}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
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

	// Check the result directly from the response body
	expected := `{"result":50}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
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

	// Check the result directly from the response body
	expected := `{"result":2}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
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

	// Check the result directly from the response body
	expected := `{"result":1}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
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

	// Check the result directly from the response body
	expected := `{"result":8}`
	actual := strings.TrimSpace(responseRecorder.Body.String()) // remove whitespace from JSON
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
