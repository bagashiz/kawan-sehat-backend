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

// UpvoteComment adds +1 vote to a comment.
func (s *Service) UpvoteComment(ctx context.Context, commentID string) (int64, error) {
	commentUUID, err := uuid.Parse(commentID)
	if err != nil {
		return 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return 0, err
	}
	_, err = s.repo.GetVoteComment(ctx, tokenPayload.AccountID, commentUUID)
	if err != nil {
		if err == ErrCommentVoteNotFound {
			return s.repo.VoteComment(ctx, tokenPayload.AccountID, commentUUID, 1)
		}
		return 0, err
	}
	return s.repo.UpdateVoteComment(ctx, tokenPayload.AccountID, commentUUID, 1)
}

// DownvoteComment reduce -1 vote to a comment.
func (s *Service) DownvoteComment(ctx context.Context, commentID string) (int64, error) {
	commentUUID, err := uuid.Parse(commentID)
	if err != nil {
		return 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return 0, err
	}
	_, err = s.repo.GetVoteComment(ctx, tokenPayload.AccountID, commentUUID)
	if err != nil {
		if err == ErrCommentVoteNotFound {
			return s.repo.VoteComment(ctx, tokenPayload.AccountID, commentUUID, -1)
		}
		return 0, err
	}
	return s.repo.UpdateVoteComment(ctx, tokenPayload.AccountID, commentUUID, -1)
}
