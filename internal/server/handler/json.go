package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// decodeAndValidateJSONRequest decodes a JSON request body into the given struct and validates it.
func (h *Handler) decodeAndValidateJSONRequest(r *http.Request, v any) error {
	if err := decodeJSONRequest(r.Body, v); err != nil {
		return BadRequest(err)
	}
	if err := h.validator.ValidateParams(v); err != nil {
		return UnprocessableRequest(err)
	}
	return nil
}

// decodeJSONRequest decodes a JSON request body into the given struct.
func decodeJSONRequest(r io.Reader, v any) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return errors.New("invalid request body")
	}
	return nil
}

// jsonResponse represents the structure of the JSON response.
type jsonResponse struct {
	Message string  `json:"message"`
	Data    any     `json:"data"`
	Error   *string `json:"error"`
}

// writeJSON sends a response in JSON format with the given status code, data, and error message.
func writeJSON(w http.ResponseWriter, statusCode int, data any, err error) error {
	msg := "success"
	var errMsg *string

	if err != nil {
		msg = "fail"
		errMsg = ptr(err.Error())
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

// ptr returns a pointer to the given string.
func ptr(s string) *string {
	return &s
}

// stringToInt32 converts a string to an int32 with a default value.
func stringToInt32(s string, i int32) (int32, error) {
	if s == "" {
		return i, nil
	}
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i64), nil
}
