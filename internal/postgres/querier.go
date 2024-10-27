// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	DeleteTopic(ctx context.Context, id uuid.UUID) (int64, error)
	InsertAccount(ctx context.Context, arg InsertAccountParams) error
	InsertTopic(ctx context.Context, arg InsertTopicParams) error
	SelectAccountByID(ctx context.Context, id uuid.UUID) (Account, error)
	SelectAccountByUsername(ctx context.Context, username string) (Account, error)
	SelectAllTopics(ctx context.Context) ([]Topic, error)
	SelectAllTopicsPaginated(ctx context.Context, arg SelectAllTopicsPaginatedParams) ([]Topic, error)
	SelectTopicByID(ctx context.Context, id uuid.UUID) (Topic, error)
	UpdateTopic(ctx context.Context, arg UpdateTopicParams) error
}

var _ Querier = (*Queries)(nil)
