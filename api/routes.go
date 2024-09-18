package api

import (
	"net/http"
)

// Set up routes for the calculator API
func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/subtract", subtractHandler)
	mux.HandleFunc("/multiply", multiplyHandler)
	mux.HandleFunc("/divide", divideHandler)
	mux.HandleFunc("/modulo", moduloHandler)
	mux.HandleFunc("/power", powerHandler)
	mux.HandleFunc("/history", historyHandler)
}
