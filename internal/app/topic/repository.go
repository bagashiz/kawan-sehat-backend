package topic

import (
	"context"

	"github.com/google/uuid"
)

// Reader is the interface that provides methods to read topic data from the storage.
type Reader interface {
	GetTopicByID(ctx context.Context, id uuid.UUID) (*Topic, error)
	ListTopics(ctx context.Context, limit, page int32) ([]*Topic, int64, error)
	ListFollowedTopics(ctx context.Context, accountID uuid.UUID, limit, page int32) ([]*FollowedTopic, int64, error)
}

// Writer is the interface that provides methods to write topic data to the storage.
type Writer interface {
	AddTopic(ctx context.Context, topic *Topic) error
	UpdateTopic(ctx context.Context, topic *Topic) error
	DeleteTopic(ctx context.Context, id uuid.UUID) error
	RelateAccountToTopic(ctx context.Context, accountTopic *AccountTopic) error
	UnrelateAccountFromTopic(ctx context.Context, accountTopic *AccountTopic) error
}

// ReadWriter is the interface that combines Reader and Writer interfaces for topic data.
type ReadWriter interface {
	Reader
	Writer
}
