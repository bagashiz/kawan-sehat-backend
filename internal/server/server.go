package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/middleware"
	"golang.org/x/sync/errgroup"
)

// Server wraps the http.Server type for extending functionality.
type server struct {
	*http.Server
}

type Config struct {
	Host string
	Port string
}

// New creates a new http.Server type, configures the routes, and adds middleware.
func New(cfg Config, h *handler.Handler, m *middleware.Middleware) *server {
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	router := registerRoutes(h, m)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &server{srv}
}

// Start starts the HTTP server in a separate goroutine and listens for
// the context cancellation signal to shut down the server gracefully.
func (s *server) Start(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)

	errs.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	errs.Go(func() error {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(), 10*time.Second,
		)
		defer cancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			return err
		}

		return nil
	})

	return errs.Wait()
}
