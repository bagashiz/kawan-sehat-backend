package server

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

// addRoutes configures the routes for the application.
func addRoutes(mux *http.ServeMux, userSvc *user.Service) {
	mux.Handle("GET /", handle(notFound()))
	mux.Handle("GET /{$}", handle(healthCheck()))

	mux.Handle("POST /v1/users/register", handle(RegisterAccount(userSvc)))
	mux.Handle("POST /v1/users/login", handle(loginAccount(userSvc)))
	mux.Handle("GET /v1/users/{id}", handle(getAccountByID(userSvc)))
}
