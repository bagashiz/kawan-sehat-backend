package server

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/validator"
)

// handlerFunc is a function that handles an HTTP request and returns an error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// handlerError is a custom error for HTTP handlers that includes a status code.
type handlerError struct {
	message    string
	statusCode int
}

// Error returns the error message for the handlerError type.
func (h *handlerError) Error() string {
	return h.message
}

func decodeAndValidateJSONRequest(r *http.Request, v any) error {
	if err := decodeJSONRequest(r.Body, v); err != nil {
		return &handlerError{message: err.Error(), statusCode: http.StatusBadRequest}
	}
	if err := validator.ValidateParams(v); err != nil {
		return &handlerError{message: err.Error(), statusCode: http.StatusUnprocessableEntity}
	}
	return nil
}

// notFound is the handler for the 404 page.
func notFound() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return &handlerError{statusCode: http.StatusNotFound, message: "404 page not found"}
	}
}

// healthCheck is the handler for the health check route.
func healthCheck() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
