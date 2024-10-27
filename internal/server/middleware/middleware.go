package middleware

import (
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
)

// Middleware holds dependencies for handling HTTP requests.
type Middleware struct {
	tokenizer user.Tokenizer
}

// New creates a new Middleware instance.
func New(tokenizer user.Tokenizer) *Middleware {
	return &Middleware{tokenizer: tokenizer}
}

// MiddlewareFunc is a function that wraps an APIFunc and returns a new APIFunc.
type MiddlewareFunc func(h handler.APIFunc) handler.APIFunc

// Chain acts as a chain of middleware functions to be applied in sequence.
func (m *Middleware) Chain(middlewares ...MiddlewareFunc) MiddlewareFunc {
	return func(next handler.APIFunc) handler.APIFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			next = middleware(next)
		}
		return next
	}
}
