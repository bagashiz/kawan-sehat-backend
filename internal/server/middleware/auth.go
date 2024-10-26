package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
)

// Auth is a middleware that checks if the request has a valid authorization header and verify if the token is valid.
// If the token is valid, the payload is added to the request context.
func Auth(h handler.APIFunc, tokenizer user.Tokenizer) handler.APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		authHeader := r.Header.Get(user.AuthHeaderKey)
		if authHeader == "" {
			return handler.UnauthorizedRequest(user.ErrAuthHeaderMissing)
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 || fields[0] != user.AuthType {
			return handler.UnauthorizedRequest(user.ErrAuthHeaderInvalid)
		}

		token := fields[1]
		payload, err := tokenizer.VerifyToken(token)
		if err != nil {
			return handler.UnauthorizedRequest(err)
		}

		ctx := context.WithValue(r.Context(), user.AuthPayloadKey, payload)
		return h(w, r.WithContext(ctx))
	}
}
