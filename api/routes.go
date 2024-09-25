package api

import (
	"net/http"
)

// Set up routes for the calculator API
func (api *API) RegisterRoutes(mux *http.ServeMux) {

	// // Public route for login
	mux.HandleFunc("/login", api.loginHandler)
	mux.HandleFunc("/register", api.registerHandler)

	// Protect routes with authentication
	mux.Handle("/add", api.authMiddleware(api.addHandler))
	mux.Handle("/subtract", api.authMiddleware(api.subtractHandler))
	mux.Handle("/multiply", api.authMiddleware(api.multiplyHandler))
	mux.Handle("/divide", api.authMiddleware(api.divideHandler))
	mux.Handle("/modulo", api.authMiddleware(api.moduloHandler))
	mux.Handle("/power", api.authMiddleware(api.powerHandler))
	mux.Handle("/history", api.authMiddleware(api.historyHandler))
	mux.Handle("/history/reset", api.authMiddleware(api.resetHandler))

}
