package reply

import (
	"time"

	"github.com/google/uuid"
)

// Reply represents a reply on a post.
type Reply struct {
	ID        uuid.UUID
	CommentID uuid.UUID
	Account   *Account
	Content   string
	Vote      *Vote
	CreatedAt time.Time
}

// Account represents a user account the reply belongs to.
type Account struct {
	ID       uuid.UUID
	Username string
}

// Vote represents a vote on a reply.
type Vote struct {
	Total int64
	State int32
}

// New creates a new comment instance.
func New(accountID, commentID, content string) (*Reply, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	postUUID, err := uuid.Parse(commentID)
	if err != nil {
		return nil, err
	}

	return &Reply{
		ID:        id,
		CommentID: postUUID,
		Account: &Account{
			ID: accountUUID,
		},
		Vote: &Vote{
			State: 0,
			Total: 0,
		},
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
