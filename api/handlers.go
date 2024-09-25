// The package api contains the handlers for the calculator API i.e. th

package api

import (
	"encoding/json"
	"net/http"
	"overengineered_calculator/storage"
)

// Generic handler for operations that return no error (Add, Subtract, Multiply, Power)
func (api *API) operationHandler(writer http.ResponseWriter, request *http.Request, operation string, functionType calculatorOperation) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := functionType(operand1, operand2)
	api.saveToHistory(operation, operand1, operand2, result)
	writeResultJSON(writer, result)
}

// Generic handler for operations that return an error (Divide, Modulo)
func (api *API) operationHandlerWithError(writer http.ResponseWriter, request *http.Request, operation string, functionType calculatorOperationWithError) {
	operand1, operand2, err := parseOperands(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := functionType(operand1, operand2)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	api.saveToHistory(operation, operand1, operand2, result)
	writeResultJSON(writer, result)
}

// Handler for Add operation
func (api *API) addHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, "Add", api.calculator.Add)
}

// Handler for Subtract operation
func (api *API) subtractHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, "Subtract", api.calculator.Subtract)
}

// Handler for Multiply operation
func (api *API) multiplyHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, "Multiply", api.calculator.Multiply)
}

// Handler for Divide operation
func (api *API) divideHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandlerWithError(writer, request, "Divide", api.calculator.Divide)
}

// Handler for Modulo operation
func (api *API) moduloHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandlerWithError(writer, request, "Modulo", api.calculator.Modulo)
}

// Handler for Power operation
func (api *API) powerHandler(writer http.ResponseWriter, request *http.Request) {
	api.operationHandler(writer, request, "Power", api.calculator.Power)
}

// Handler for retrieving history
func (api *API) historyHandler(writer http.ResponseWriter, request *http.Request) {
	history, err := api.storage.GetHistory()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(history)
}

// Handler for resetting calculator history
func (api *API) resetHandler(writer http.ResponseWriter, request *http.Request) {
	api.storage.ResetHistory()
	writer.WriteHeader(http.StatusOK)
}

// Handler for user login.
func (api *API) loginHandler(writer http.ResponseWriter, request *http.Request) {

	// Decode the user credentials from the request body
	var user storage.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Invalid credentials format", http.StatusBadRequest)
		return
	}

	// Authenticate user using storage strategy
	err = api.storage.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Generate JWT (JSON Web Token)
	token, err := generateJWT(user.Username)
	if err != nil {
		http.Error(writer, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// // Set the token as a cookie in the response
	// http.SetCookie(writer, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   token,
	// 	Expires: time.Now().Add(24 * time.Hour),
	// })
	// json.NewEncoder(writer).Encode(map[string]string{"token": token})

	// Return the token in the response body
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{
		"token": "Bearer " + token, // Adding 'Bearer' as per jwt.io/introduction
	})
}

// Handler for user registration
func (api *API) registerHandler(writer http.ResponseWriter, request *http.Request) {

	// Decode the user credentials from the request body
	var user storage.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Use storage strategy to register the user
	err = api.storage.RegisterUser(user.Username, user.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			http.Error(writer, "User already exists", http.StatusConflict)
		} else {
			http.Error(writer, "Could not register user", http.StatusInternalServerError)
		}
		return
	}

	// Return success message
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(map[string]string{"message": "User registered successfully"})
}
