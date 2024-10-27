package handler

import "net/http"

// APIError is a custom error for HTTP handlers that includes a status code.
type APIError struct {
	Message    string
	StatusCode int
}

// Error returns the error message for the APIError type.
func (h APIError) Error() string {
	return h.Message
}

// BadRequest returns a HandlerError for an invalid request body with 400 status code.
func BadRequest(err error) APIError {
	return APIError{Message: err.Error(), StatusCode: http.StatusBadRequest}
}

// UnauthorizedRequest returns a HandlerError for an UnauthorizedRequest request with 401 status code.
func UnauthorizedRequest(err error) APIError {
	return APIError{Message: err.Error(), StatusCode: http.StatusUnauthorized}
}

// ForbiddenRequest returns a HandlerError for a forbidden request with 403 status code.
func ForbiddenRequest(err error) APIError {
	return APIError{Message: err.Error(), StatusCode: http.StatusForbidden}
}

// NotFoundRequest returns a HandlerError for a not found request with 404 status code.
func NotFoundRequest(err error) APIError {
	return APIError{Message: err.Error(), StatusCode: http.StatusNotFound}
}

// ConflictRequest returns a HandlerError for a ConflictRequest request with 409 status code.
func ConflictRequest(err error) APIError {
	return APIError{Message: err.Error(), StatusCode: http.StatusConflict}
}

// UnprocessableRequest returns a HandlerError for an unprocessable request with 422 status code.
func UnprocessableRequest(err error) APIError {
	return APIError{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
}
