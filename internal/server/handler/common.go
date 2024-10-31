package handler

import (
	"errors"
	"net/http"
)

// NotFound is the handler for the 404 page.
func (h *Handler) NotFound() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := errors.New("404 page not found")
		return NotFoundRequest(err)
	}
}

// NotAllowed is the handler for the 405 page.
func (h *Handler) NotAllowed() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := errors.New("405 method not allowed")
		return NotAllowedRequest(err)
	}
}

// HealthCheck is the handler for the health check route.
func (h *Handler) HealthCheck() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
