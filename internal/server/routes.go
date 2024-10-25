package server

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

func addRoutes(mux *http.ServeMux, userSvc *user.Service) {
	mux.Handle("GET /", handle(notFound()))
	mux.Handle("GET /{$}", handle(healthCheck()))

	mux.Handle("POST /v1/users/register", handle(RegisterAccount(userSvc)))
	mux.Handle("POST /v1/users/login", handle(loginAccount(userSvc)))
}
