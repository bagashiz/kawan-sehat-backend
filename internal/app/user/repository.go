package user

import (
	"context"

	"github.com/google/uuid"
)

// Reader is the interface that provides methods to read user data from the storage.
type Reader interface {
	GetAccountByUsername(ctx context.Context, username string) (*Account, error)
	GetAccountByID(ctx context.Context, id uuid.UUID) (*Account, error)
	ListIllnessHistoriesByAccountID(ctx context.Context, id uuid.UUID) ([]*IllnessHistory, error)
}

// Writer is the interface that provides methods to write user data to the storage.
type Writer interface {
	AddAccount(ctx context.Context, account *Account) error
	UpdateAccount(ctx context.Context, account *Account) error
}

// ReadWriter is the interface that combines Reader and Writer interfaces for user data.
type ReadWriter interface {
	Reader
	Writer
}
