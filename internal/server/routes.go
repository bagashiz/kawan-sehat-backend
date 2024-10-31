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
	replyRouter := replyRoutes(h, m)

	mux.Route("/v1", func(r chi.Router) {
		r.Mount("/users", userRouter)
		r.Mount("/topics", topicRouter)
		r.Mount("/posts", postRouter)
		r.Mount("/comments", commentRouter)
		r.Mount("/replies", replyRouter)
	})

	mux.NotFound(handle(h.NotFound()))
	mux.MethodNotAllowed(handle(h.NotAllowed()))
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
		mux.Get("/illnesses", handle(auth(h.ListAccountIllnessHistories())))
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
		mux.Post("/{id}/upvote", handle(auth(h.UpvotePost())))
		mux.Post("/{id}/downvote", handle(auth(h.DownvotePost())))
		mux.Put("/{id}", handle(auth(h.UpdatePost())))
		mux.Delete("/{id}", handle(auth(h.DeletePost())))
		mux.Delete("/{id}/unmark", handle(auth(h.Unbookmark())))
		mux.Get("/{id}", handle(auth(h.GetPostByID())))
		mux.Get("/", handle(auth(h.ListPosts())))
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
		mux.Post("/{id}/upvote", handle(auth(h.UpvoteComment())))
		mux.Post("/{id}/downvote", handle(auth(h.DownvoteComment())))
		mux.Delete("/{id}", handle(auth(h.DeleteComment())))
		mux.Get("/{id}/replies", handle(auth(h.ListRepliesByCommentID())))
	})
	return mux
}

// replyRoutes configures the routes for the post service.
func replyRoutes(h *handler.Handler, m *middleware.Middleware) *chi.Mux {
	auth := m.Auth
	mux := chi.NewRouter()
	mux.Route("/replies", func(r chi.Router) {
		mux.Post("/", handle(auth(h.CreateReply())))
		mux.Post("/{id}/upvote", handle(auth(h.UpvoteReply())))
		mux.Post("/{id}/downvote", handle(auth(h.DownvoteReply())))
		mux.Delete("/{id}", handle(auth(h.DeleteReply())))
	})
	return mux
}

// handle wraps the handler.Handle function to shorten the function signature.
func handle(h handler.APIFunc) http.HandlerFunc {
	return handler.Handle(h)
}
