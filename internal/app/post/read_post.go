package post

import (
	"context"

	"github.com/google/uuid"
)

// GetPost retrieves post data from the repository by ID.
func (s *Service) GetPost(ctx context.Context, id string) (*Post, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repo.GetPostByID(ctx, uuid)
}

// ListPostsParams defines the parameters to list posts.
type ListPostsParams struct {
	AccountID string
	TopicID   string
	Limit     int32
	Offset    int32
}

// ListPosts lists all posts from the repository with optional filters.
func (s *Service) ListPosts(ctx context.Context, params ListPostsParams) ([]*Post, error) {
	if params.AccountID != "" {
		uuid, err := uuid.Parse(params.AccountID)
		if err != nil {
			return nil, err
		}
		return s.repo.ListPostsByAccountID(ctx, uuid, params.Limit, params.Offset)
	}
	if params.TopicID != "" {
		uuid, err := uuid.Parse(params.TopicID)
		if err != nil {
			return nil, err
		}
		return s.repo.ListPostsByTopicID(ctx, uuid, params.Limit, params.Offset)
	}
	return s.repo.ListPosts(ctx, params.Limit, params.Offset)
}

// ListPostsByTopic lists all posts from the repository by topic ID with optional filters.
func (s *Service) ListPostsByTopic(ctx context.Context, topicID string, limit, offset int32) ([]*Post, error) {
	uuid, err := uuid.Parse(topicID)
	if err != nil {
		return nil, err
	}
	return s.repo.ListPostsByTopicID(ctx, uuid, limit, offset)
}

// ListPostsByAccount lists all posts from the repository by account ID with optional filters.
func (s *Service) ListPostsByAccount(ctx context.Context, accountID string, limit, offset int32) ([]*Post, error) {
	uuid, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}
	return s.repo.ListPostsByAccountID(ctx, uuid, limit, offset)
}
