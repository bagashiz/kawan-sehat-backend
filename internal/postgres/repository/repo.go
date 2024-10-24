package repository

import (
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
)

// PostgresRepository is a repository that holds the database connection.
type PostgresRepository struct {
	db *postgres.DB
}

// New creates a new PostgresRepository instance.
func New(db *postgres.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
