package server

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

// registerRoutes configures the routes for the application.
func registerRoutes(h *handler.Handler, m *middleware.Middleware) *chi.Mux {
	mux := chi.NewRouter()
	userRouter := userRoutes(h, m)
	topicRouter := topicRoutes(h, m)

	mux.Route("/v1", func(r chi.Router) {
		r.Mount("/users", userRouter)
		r.Mount("/topics", topicRouter)
	})

	mux.Get("/", handle(h.NotFound()))
	mux.Get("/healthz", handle(h.HealthCheck()))

	return mux
}

// userRoutes configures the routes for the user service.
func userRoutes(h *handler.Handler, m *middleware.Middleware) *chi.Mux {
	auth := m.Auth
	mux := chi.NewRouter()

	mux.Route("/users", func(r chi.Router) {
		mux.Post("/register", handle(h.RegisterAccount()))
		mux.Post("/login", handle(h.LoginAccount()))
		mux.Put("/", handle(auth(h.UpdateAccount())))
		mux.Get("/{id}", handle(h.GetAccountByID()))
		mux.Get("/topics", handle(auth(h.ListFollowedTopics())))
	})
	return mux
}

// topicRoutes configures the routes for the topic service.
func topicRoutes(h *handler.Handler, m *middleware.Middleware) *chi.Mux {
	auth := m.Auth
	admin := m.Chain(m.Auth, m.Admin)
	mux := chi.NewRouter()

	mux.Route("/topics", func(r chi.Router) {
		mux.Post("/", handle(admin(h.CreateTopic())))
		mux.Post("/{id}/follow", handle(auth(h.FollowTopic())))
		mux.Put("/{id}", handle(admin(h.UpdateTopic())))
		mux.Delete("/{id}/unfollow", handle(auth(h.UnfollowTopic())))
		mux.Delete("/{id}", handle(admin(h.DeleteTopic())))
		mux.Get("/{id}", handle(h.GetTopicByID()))
		mux.Get("/", handle(h.ListTopics()))
	})
	return mux
}

// handle wraps the handler.Handle function to shorten the function signature.
func handle(h handler.APIFunc) http.HandlerFunc {
	return handler.Handle(h)
}
