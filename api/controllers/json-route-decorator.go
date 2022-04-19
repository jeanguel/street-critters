package controllers

import (
	"encoding/json"
	"net/http"
)

// JSONRoute
func JSONRoute(f func(w http.ResponseWriter, r *http.Request) (int, interface{})) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// CORS-related headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		status, response := f(w, r)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
	}
}
