package middleware

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
)

// Admin is a middleware that checks if the current request is made by a user with admin role.
func (m *Middleware) Admin(h handler.APIFunc) handler.APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		payload, err := user.GetTokenPayload(r.Context())
		if err != nil {
			return handler.UnauthorizedRequest(err)
		}

		if payload.UserRole != user.Admin {
			return handler.ForbiddenRequest(user.ErrAccountForbidden)
		}

		return h(w, r)
	}
}
