package server

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
	"github.com/go-chi/chi/v5"
)

// registerRoutes configures the routes for the application.
func registerRoutes(userSvc *user.Service) *chi.Mux {
	mux := chi.NewRouter()

	mux.Route("/users", func(r chi.Router) {
		mux.Post("/register", handle(handler.RegisterAccount(userSvc)))
		mux.Post("/login", handle(handler.LoginAccount(userSvc)))
		mux.Get("/{id}", handle(handler.GetAccountByID(userSvc)))
	})

	mux.Route("/v1", func(r chi.Router) {
		r.Mount("/users", mux)
	})

	mux.Get("/", handle(handler.NotFound()))
	mux.Get("/healthz", handle(handler.HealthCheck()))

	return mux
}

// handle wraps the handler.Handle function to shorten the function signature.
func handle(h handler.APIFunc) http.HandlerFunc {
	return handler.Handle(h)
}
