package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/reply"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/validator"
)

// Handler holds dependencies for handling HTTP requests.
type Handler struct {
	validator  *validator.Validator
	userSvc    *user.Service
	topicSvc   *topic.Service
	postSvc    *post.Service
	commentSvc *comment.Service
	replySvc   *reply.Service
}

// New creates a new Handler instance.
func New(
	validator *validator.Validator,
	userSvc *user.Service,
	topicSvc *topic.Service,
	postSvc *post.Service,
	commentSvc *comment.Service,
	replySvc *reply.Service,
) *Handler {
	return &Handler{
		validator:  validator,
		userSvc:    userSvc,
		topicSvc:   topicSvc,
		postSvc:    postSvc,
		commentSvc: commentSvc,
		replySvc:   replySvc,
	}
}

// APIFunc is a function that handles an HTTP request and returns an error.
type APIFunc func(http.ResponseWriter, *http.Request) error

// responseWriter extends the http.ResponseWriter type to store the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the WriteHeader method to store the status code.
func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Handle wraps custom APIFunc type as an http.HandlerFunc,
// switch to custom ResponseWriter, handles errors, and logs request information.
func Handle(h APIFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "*")

		if err := h(writer, r); err != nil {
			statusCode := http.StatusInternalServerError
			errMsg := http.StatusText(statusCode)

			if apiErr, ok := err.(APIError); ok {
				statusCode = apiErr.StatusCode
				errMsg = apiErr.Error()
			}

			if encodeErr := writeJSON(writer, statusCode, nil, err); encodeErr != nil {
				http.Error(writer, errMsg, http.StatusInternalServerError)
			}

			logRequest(r, statusCode, time.Since(start), errMsg)
			return
		}

		logRequest(r, writer.statusCode, time.Since(start), "")
	})
}

// logRequest logs request information based on the status code and error message.
func logRequest(r *http.Request, statusCode int, duration time.Duration, errMsg string) {
	switch {
	case statusCode < 400:
		slog.Info("REQUEST_SUCCESS",
			slog.Int("status", statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", duration),
		)
	case statusCode >= 400 && statusCode < 500:
		slog.Warn("CLIENT_ERROR",
			slog.Int("status", statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", duration),
		)
	case statusCode >= 500 || errMsg != "":
		slog.Error("SERVER_ERROR",
			slog.Int("status", statusCode),
			slog.String("error", errMsg),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", duration),
		)
	}
}
