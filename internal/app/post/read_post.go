package post

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// GetPost retrieves post data from the repository by ID.
func (s *Service) GetPost(ctx context.Context, id string) (*Post, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	return s.repo.GetPostByID(ctx, tokenPayload.AccountID, uuid)
}

// ListPostsParams defines the parameters to list posts.
type ListPostsParams struct {
	AccountID string
	TopicID   string
	Limit     int32
	Page      int32
}

// ListPosts lists all posts from the repository with optional filters.
func (s *Service) ListPosts(ctx context.Context, params ListPostsParams) ([]*Post, int64, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, 0, err
	}
	accountID := tokenPayload.AccountID
	if params.AccountID != "" {
		uuid, err := uuid.Parse(params.AccountID)
		if err != nil {
			return nil, 0, err
		}
		return s.repo.ListPostsByAccountID(ctx, uuid, params.Limit, params.Page)
	}
	if params.TopicID != "" {
		uuid, err := uuid.Parse(params.TopicID)
		if err != nil {
			return nil, 0, err
		}
		return s.repo.ListPostsByTopicID(ctx, accountID, uuid, params.Limit, params.Page)
	}
	return s.repo.ListPosts(ctx, accountID, params.Limit, params.Page)
}

// ListPostsByTopic lists all posts from the repository by topic ID with optional filters.
func (s *Service) ListPostsByTopic(ctx context.Context, topicID string, limit, page int32) ([]*Post, int64, error) {
	uuid, err := uuid.Parse(topicID)
	if err != nil {
		return nil, 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, 0, err
	}
	accountID := tokenPayload.AccountID
	return s.repo.ListPostsByTopicID(ctx, accountID, uuid, limit, page)
}

// ListPostsByAccount lists all posts from the repository by account ID with optional filters.
func (s *Service) ListPostsByAccount(ctx context.Context, accountID string, limit, page int32) ([]*Post, int64, error) {
	uuid, err := uuid.Parse(accountID)
	if err != nil {
		return nil, 0, err
	}
	return s.repo.ListPostsByAccountID(ctx, uuid, limit, page)
}

// ListBookmarks lists all bookmarked posts from the repository by account ID with optional filters.
func (s *Service) ListBookmarks(ctx context.Context, limit, page int32) ([]*Post, int64, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, 0, err
	}
	accountID := tokenPayload.AccountID
	return s.repo.ListAccountBookmarks(ctx, accountID, limit, page)
}
