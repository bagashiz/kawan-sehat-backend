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
	postRouter := postRoutes(h, m)
	commentRouter := commentRoutes(h, m)

	mux.Route("/v1", func(r chi.Router) {
		r.Mount("/users", userRouter)
		r.Mount("/topics", topicRouter)
		r.Mount("/posts", postRouter)
		r.Mount("/comments", commentRouter)
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
		mux.Get("/bookmarks", handle(auth(h.ListBookmarks())))
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

// postRoutes configures the routes for the post service.
func postRoutes(h *handler.Handler, m *middleware.Middleware) *chi.Mux {
	auth := m.Auth
	mux := chi.NewRouter()
	mux.Route("/posts", func(r chi.Router) {
		mux.Post("/", handle(auth(h.CreatePost())))
		mux.Post("/{id}/mark", handle(auth(h.Bookmark())))
		mux.Put("/{id}", handle(auth(h.UpdatePost())))
		mux.Delete("/{id}", handle(auth(h.DeletePost())))
		mux.Delete("/{id}/unmark", handle(auth(h.Unbookmark())))
		mux.Get("/{id}", handle(h.GetPostByID()))
		mux.Get("/", handle(h.ListPosts()))
		mux.Get("/{id}/comments", handle(auth(h.ListCommentsByPostID())))
	})
	return mux
}

// commentRoutes configures the routes for the post service.
func commentRoutes(h *handler.Handler, m *middleware.Middleware) *chi.Mux {
	auth := m.Auth
	mux := chi.NewRouter()
	mux.Route("/comments", func(r chi.Router) {
		mux.Post("/", handle(auth(h.CreateComment())))
		mux.Delete("/{id}", handle(auth(h.DeleteComment())))
	})
	return mux
}

// handle wraps the handler.Handle function to shorten the function signature.
func handle(h handler.APIFunc) http.HandlerFunc {
	return handler.Handle(h)
}
