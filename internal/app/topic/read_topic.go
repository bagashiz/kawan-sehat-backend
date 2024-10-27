package topic

import (
	"context"

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
