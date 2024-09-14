package api

import (
	"net/http"
)

// Set up routes for the calculator API
func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/add", AddHandler)
	mux.HandleFunc("/subtract", SubtractHandler)
	mux.HandleFunc("/multiply", MultiplyHandler)
	mux.HandleFunc("/divide", DivideHandler)
	mux.HandleFunc("/modulo", ModuloHandler)
	mux.HandleFunc("/power", PowerHandler)
	mux.HandleFunc("/history", HistoryHandler)
}
