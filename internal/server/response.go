package server

import (
	"encoding/json"
	"net/http"
)

// JSONResponse represents the structure of the JSON response.
type jsonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// sendJSONResponse sends a response in JSON format with the given status code, data, and error message.
func sendJSONResponse(w http.ResponseWriter, statusCode int, data any, err error) error {
	var msg, errMsg string

	if err != nil {
		msg = "fail"
		errMsg = err.Error()
	} else {
		msg = "success"
	}

	response := &jsonResponse{
		Message: msg,
		Data:    data,
		Error:   errMsg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(response)
}
