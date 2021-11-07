package restapi

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, error interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(error)
}

func (errorResponse *errorResponse) message (message string) *errorResponse{
	errorResponse.Message = message
	return errorResponse
}