package topic

import (
	"context"

	"github.com/google/uuid"
)

// CreateTopicParams defines the parameters to create a new topic.
type CreateTopicParams struct {
	Name        string
	Description string
}

// AddTopic creates a new topic and stores it in the repository.
func (s *Service) AddTopic(ctx context.Context, params CreateTopicParams) (*Topic, error) {
	topic, err := New(params.Name, params.Description)
	if err != nil {
		return nil, err
	}

	if err := s.repo.AddTopic(ctx, topic); err != nil {
		return nil, err
	}

	return topic, nil
}

// UpdateTopicParams defines the parameters to update an existing topic.
type UpdateTopicParams struct {
	ID          string
	Name        string
	Description string
}

// UpdateTopic updates an existing topic in the repository.
func (s *Service) UpdateTopic(ctx context.Context, params UpdateTopicParams) (*Topic, error) {
	uuid, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	topic, err := s.repo.GetTopicByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	topic.Update(params.Name, params.Description)

	if err := s.repo.UpdateTopic(ctx, topic); err != nil {
		return nil, err
	}

	return topic, nil
}

// DeleteTopic deletes a topic from the repository by its ID.
func (s *Service) DeleteTopic(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteTopic(ctx, uuid)
}
