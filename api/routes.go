package api

import (
	"net/http"
)

// Set up routes for the calculator API
func (api *API) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/add", api.addHandler)
	mux.HandleFunc("/subtract", api.subtractHandler)
	mux.HandleFunc("/multiply", api.multiplyHandler)
	mux.HandleFunc("/divide", api.divideHandler)
	mux.HandleFunc("/modulo", api.moduloHandler)
	mux.HandleFunc("/power", api.powerHandler)
	mux.HandleFunc("/history", api.historyHandler)
	mux.HandleFunc("/reset", api.resetHandler)

}
