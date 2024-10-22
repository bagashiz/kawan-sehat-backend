package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// DB wraps the Querier interface and a pgxpool.Pool instance.
type DB struct {
	Querier
	*pgxpool.Pool
}

// Connect creates a new DB instance using the provided config.
func Connect(ctx context.Context, connURI string) (*DB, error) {
	pool, err := pgxpool.New(ctx, connURI)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{New(pool), pool}, nil
}

// Migrate runs the goose migration tool to apply new migrations.
func (d *DB) Migrate() error {
	goose.SetBaseFS(migrationFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(d.Pool)
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}

// ExecTX wraps the provided function in a transaction and executes it.
func (d *DB) ExecTX(ctx context.Context, fn func(Querier) error) error {
	tx, err := d.Begin(ctx)
	if err != nil {
		return err
	}

	if err := fn(New(tx)); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit(ctx)
}
