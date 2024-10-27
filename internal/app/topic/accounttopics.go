package topic

import (
	"context"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// AccountTopic represents a relationship between an account and a topic.
type AccountTopic struct {
	AccountID uuid.UUID
	TopicID   uuid.UUID
	CreatedAt time.Time
}

// FollowedTopic represents a topic that an account follows.
type FollowedTopic struct {
	*Topic
}

// NewAccountTopic creates a new AccountTopic instance.
func NewAccountTopic(accountID, topicID string) (*AccountTopic, error) {
	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	topicUUID, err := uuid.Parse(topicID)
	if err != nil {
		return nil, err
	}

	return &AccountTopic{
		AccountID: accountUUID,
		TopicID:   topicUUID,
		CreatedAt: time.Now(),
	}, nil
}

// createAccountTopic is a helper function to create new AccountTopic instance from the context.
func createAccountTopic(ctx context.Context, topicID string) (*AccountTopic, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}

	accountID := tokenPayload.AccountID.String()
	accountTopic, err := NewAccountTopic(accountID, topicID)
	if err != nil {
		return nil, err
	}

	return accountTopic, nil
}
