package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

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

	tknzr, err := token.New(cfg.Token)
	if err != nil {
		return err
	}

	db, err := postgres.Connect(ctx, cfg.DB.URI)
	if err != nil {
		return err
	}
	defer db.Close()

	slog.Info("connected to the database", "type", cfg.DB.Type)

	if err := db.Migrate(); err != nil {
		return err
	}

	vldtr := validator.New()
	pgRepo := repository.New(db)

	userSvc := user.NewService(pgRepo, tknzr)
	topicSvc := topic.NewService(pgRepo)

	hndlr := handler.New(vldtr, userSvc, topicSvc)
	mw := middleware.New(tknzr)
	srv := server.New(cfg.App, hndlr, mw)

	slog.Info("starting the http server", "addr", srv.Addr)

	if err := srv.Start(ctx); err != nil {
		return err
	}

	return nil
}
