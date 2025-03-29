package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type MigrationDirection string

const (
	MigrationDirectionUp   MigrationDirection = "up"
	MigrationDirectionDown MigrationDirection = "down"
)

// Migrate runs the goose migration tool to apply new migrations.
func Migrate(ctx context.Context, uri string, direction MigrationDirection) error {
	goose.SetBaseFS(migrationFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error setting dialect: %w", err)
	}

	dbCfg, err := pgx.ParseConfig(uri)
	if err != nil {
		return fmt.Errorf("error parsing connection config: %w", err)
	}

	db := stdlib.OpenDB(*dbCfg)
	defer db.Close()

	switch direction {
	case MigrationDirectionUp:
		if err := goose.Up(db, "migrations"); err != nil {
			return err
		}
	case MigrationDirectionDown:
		if err := goose.Down(db, "migrations"); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	return nil
}
