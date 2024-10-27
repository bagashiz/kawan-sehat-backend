package server

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

// registerRoutes configures the routes for the application.
func registerRoutes(t user.Tokenizer, userSvc *user.Service, topicSvc *topic.Service) *chi.Mux {
	mux := chi.NewRouter()
	userRouter := userRoutes(userSvc)
	topicRouter := topicRoutes(t, topicSvc)

	mux.Route("/v1", func(r chi.Router) {
		r.Mount("/users", userRouter)
		r.Mount("/topics", topicRouter)
	})

	mux.Get("/", handle(handler.NotFound()))
	mux.Get("/healthz", handle(handler.HealthCheck()))

	return mux
}

func userRoutes(userSvc *user.Service) *chi.Mux {
	mux := chi.NewRouter()
	mux.Route("/users", func(r chi.Router) {
		mux.Post("/register", handle(handler.RegisterAccount(userSvc)))
		mux.Post("/login", handle(handler.LoginAccount(userSvc)))
		mux.Get("/{id}", handle(handler.GetAccountByID(userSvc)))
	})
	return mux
}

func topicRoutes(t user.Tokenizer, topicSvc *topic.Service) *chi.Mux {
	mux := chi.NewRouter()
	mux.Route("/topics", func(r chi.Router) {
		mux.Post("/",
			handle(auth(admin(handler.CreateTopic(topicSvc)), t)),
		)
		mux.Put("/{id}",
			handle(auth(admin(handler.UpdateTopic(topicSvc)), t)),
		)
		mux.Delete("/{id}",
			handle(auth(admin(handler.DeleteTopic(topicSvc)), t)),
		)
		mux.Get("/{id}", handle(handler.GetTopicByID(topicSvc)))
		mux.Get("/", handle(handler.ListTopics(topicSvc)))
	})
	return mux
}

// handle wraps the handler.Handle function to shorten the function signature.
func handle(h handler.APIFunc) http.HandlerFunc {
	return handler.Handle(h)
}

// admin wraps the middleware.Admih function to shorten the function signature.
func admin(h handler.APIFunc) handler.APIFunc {
	return middleware.Admin(h)
}

// auth wraps the middleware.Admih function to shorten the function signature.
func auth(h handler.APIFunc, t user.Tokenizer) handler.APIFunc {
	return middleware.Auth(h, t)
}
