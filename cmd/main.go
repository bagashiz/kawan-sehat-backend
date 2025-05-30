package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/reply"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/config"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres/repository"
	"github.com/bagashiz/kawan-sehat-backend/internal/server"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/handler"
	"github.com/bagashiz/kawan-sehat-backend/internal/server/middleware"
	"github.com/bagashiz/kawan-sehat-backend/internal/token"
	"github.com/bagashiz/kawan-sehat-backend/internal/validator"
)

// entry point of the application.
func main() {
	ctx := context.Background()

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)
	slog.SetDefault(logger)

	if err := run(ctx, os.Getenv); err != nil {
		slog.Error("error running application", "error", err)
		os.Exit(1)
	}
}

// run sets up dependencies and starts the application.
func run(ctx context.Context, getEnv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg, err := config.New(getEnv)
	if err != nil {
		return err
	}

	tknzr, err := token.New(cfg.Token().Type, cfg.Token().Secret)
	if err != nil {
		return err
	}

	db, err := postgres.NewDB(ctx, cfg.DB().URI)
	if err != nil {
		return err
	}
	defer db.Close()

	slog.Info("connected to the database", "type", cfg.DB().Type)

	if err := postgres.Migrate(ctx, cfg.DB().URI, postgres.MigrationDirectionUp); err != nil {
		return err
	}

	vldtr := validator.New()
	pgRepo := repository.New(db)

	userSvc := user.NewService(pgRepo, tknzr)
	topicSvc := topic.NewService(pgRepo)
	postSvc := post.NewService(pgRepo)
	commentSvc := comment.NewService(pgRepo)
	replySvc := reply.NewService(pgRepo)

	hndlr := handler.New(
		vldtr, userSvc, topicSvc, postSvc, commentSvc, replySvc,
	)
	mw := middleware.New(tknzr)
	srv := server.New(server.Config{
		Host: cfg.App().Host,
		Port: cfg.App().Port,
	}, hndlr, mw)

	slog.Info("starting the http server", "addr", srv.Addr)

	if err := srv.Start(ctx); err != nil {
		return err
	}

	return nil
}
