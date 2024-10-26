package handler

import (
	"errors"
	"net/http"
)

// NotFound is the handler for the 404 page.
func NotFound() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := errors.New("404 page not found")
		return NotFoundRequest(err)
	}
}

// HealthCheck is the handler for the health check route.
func HealthCheck() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
