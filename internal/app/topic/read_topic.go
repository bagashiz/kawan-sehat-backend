package topic

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// GetTopicByID gets a topic from the repository by its ID.
func (s *Service) GetTopicByID(ctx context.Context, id string) (*Topic, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return s.repo.GetTopicByID(ctx, uuid)
}

// ListTopics lists all topics from the repository with optional filters.
func (s *Service) ListTopics(ctx context.Context, limit, offset int32) ([]*Topic, error) {
	return s.repo.ListTopics(ctx, limit, offset)
}

// ListFollowedTopics list all account's followed topics with optional filters.
func (s *Service) ListFollowedTopics(ctx context.Context, limit, offset int32) ([]*FollowedTopic, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	accountID := tokenPayload.AccountID
	return s.repo.ListFollowedTopics(ctx, accountID, limit, offset)
}
