package comment

import (
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// Comment represents a comment on a post.
type Comment struct {
	ID           uuid.UUID
	PostID       uuid.UUID
	Account      *Account
	Content      string
	Vote         *Vote
	TotalReplies int64
	CreatedAt    time.Time
}

// Account represents a user account the comment belongs to.
type Account struct {
	ID       uuid.UUID
	Username string
	Avatar   user.Avatar
	Role     user.Role
}

// Vote represents a vote on a comment.
type Vote struct {
	Total int64
	State int32
}

// New creates a new comment instance.
func New(accountID, postID, content string) (*Comment, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	return &Comment{
		ID:     id,
		PostID: postUUID,
		Account: &Account{
			ID: accountUUID,
		},
		Vote: &Vote{
			State: 0,
			Total: 0,
		},
		TotalReplies: 0,
		Content:      content,
		CreatedAt:    time.Now(),
	}, nil
}
