package post

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// CreatePostParams defines the parameters to create a new post.
type CreatePostParams struct {
	TopicID string
	Title   string
	Content string
}

// AddPost creates a new post and stores it in the repository.
func (s *Service) AddPost(ctx context.Context, params CreatePostParams) (*Post, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	accountID := tokenPayload.AccountID.String()
	post, err := New(
		accountID, params.TopicID, params.Title, params.Content,
	)
	if err != nil {
		return nil, err
	}
	if err := s.repo.AddPost(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

// UpdatePostParams defines the parameters to update an existing post.
type UpdatePostParams struct {
	ID      string
	Title   string
	Content string
}

// UpdatePost updates an existing post in the repository.
func (s *Service) UpdatePost(ctx context.Context, params UpdatePostParams) (*Post, error) {
	uuid, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	post, err := s.repo.GetPostByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if tokenPayload.AccountID != post.AccountID {
		return nil, ErrPostActionForbidden
	}
	post.Update(params.Title, params.Content)
	if err := s.repo.UpdatePost(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePost deletes a post from the repository by its ID.
func (s *Service) DeletePost(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return err
	}
	post, err := s.repo.GetPostByID(ctx, uuid)
	if err != nil {
		return err
	}
	if tokenPayload.AccountID != post.AccountID && tokenPayload.AccountRole != user.Admin {
		return ErrPostActionForbidden
	}
	return s.repo.DeletePost(ctx, uuid)
}
