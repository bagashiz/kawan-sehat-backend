package server

import "net/http"

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

// notFound is the handler for the 404 page.
func notFound() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return &handlerError{statusCode: http.StatusNotFound, message: "route not found"}
	}
}

// index is the handler for the landing page.
func index() handlerFunc {
	type indexResponse struct {
		Message string `json:"message"`
	}
	return func(w http.ResponseWriter, r *http.Request) error {
		return sendJSONResponse(w, http.StatusOK, &indexResponse{"hello world"}, nil)
	}
}
