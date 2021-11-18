package controller

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func (errorResponse *errorResponse) message(message string) *errorResponse {
	errorResponse.Message = message
	return errorResponse
}
