package comment

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// CreateCommentParams defines the parameters to create a new comment.
type CreateCommentParams struct {
	PostID  string
	Content string
}

// AddComment adds a new comment.
func (s *Service) AddComment(ctx context.Context, params CreateCommentParams) (*Comment, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	accountID := tokenPayload.AccountID.String()
	comment, err := New(accountID, params.PostID, params.Content)
	if err != nil {
		return nil, err
	}
	if err := s.repo.AddComment(ctx, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// DeleteComment deletes a comment by its ID.
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return err
	}
	comment, err := s.repo.GetCommentByID(ctx, uuid)
	if err != nil {
		return err
	}
	if tokenPayload.AccountID != comment.Account.ID && tokenPayload.AccountRole != user.Admin {
		return ErrCommentActionForbidden
	}
	return s.repo.DeleteComment(ctx, uuid)
}
