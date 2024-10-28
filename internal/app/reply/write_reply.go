package reply

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// CreateReplyParams defines the parameters to create a new reply.
type CreateReplyParams struct {
	CommentID string
	Content   string
}

// AddReply adds a new reply.
func (s *Service) AddReply(ctx context.Context, params CreateReplyParams) (*Reply, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	accountID := tokenPayload.AccountID.String()
	reply, err := New(accountID, params.CommentID, params.Content)
	if err != nil {
		return nil, err
	}
	if err := s.repo.AddReply(ctx, reply); err != nil {
		return nil, err
	}
	return reply, nil
}

// DeleteReply deletes a reply by its ID.
func (s *Service) DeleteReply(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return err
	}
	reply, err := s.repo.GetReplyByID(ctx, uuid)
	if err != nil {
		return err
	}
	if tokenPayload.AccountID != reply.Account.ID && tokenPayload.AccountRole != user.Admin {
		return ErrReplyActionForbidden
	}
	return s.repo.DeleteReply(ctx, uuid)
}

// UpvoteReply adds +1 vote to a reply.
func (s *Service) UpvoteReply(ctx context.Context, replyID string) (int64, error) {
	replyUUID, err := uuid.Parse(replyID)
	if err != nil {
		return 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return 0, err
	}
	_, err = s.repo.GetVoteReply(ctx, tokenPayload.AccountID, replyUUID)
	if err != nil {
		if err == ErrReplyVoteNotFound {
			return s.repo.VoteReply(ctx, tokenPayload.AccountID, replyUUID, 1)
		}
		return 0, err
	}
	return s.repo.UpdateVoteReply(ctx, tokenPayload.AccountID, replyUUID, 1)
}

// DownvoteReply adds +1 vote to a reply.
func (s *Service) DownvoteReply(ctx context.Context, replyID string) (int64, error) {
	replyUUID, err := uuid.Parse(replyID)
	if err != nil {
		return 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return 0, err
	}
	_, err = s.repo.GetVoteReply(ctx, tokenPayload.AccountID, replyUUID)
	if err != nil {
		if err == ErrReplyVoteNotFound {
			return s.repo.VoteReply(ctx, tokenPayload.AccountID, replyUUID, -1)
		}
		return 0, err
	}
	return s.repo.UpdateVoteReply(ctx, tokenPayload.AccountID, replyUUID, -1)
}
