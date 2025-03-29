package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB wraps the [Querier] interface and a [pgxpool.Pool] instance.
// Extends functionality for migrations and transactions.
type DB struct {
	Querier
	*pgxpool.Pool
}

// NewDB creates a new DB instance using the provided config.
func NewDB(ctx context.Context, connURI string) (*DB, error) {
	pool, err := pgxpool.New(ctx, connURI)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{New(pool), pool}, nil
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
