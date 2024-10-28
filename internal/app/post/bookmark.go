package post

import (
	"context"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// Bookmark represents a relationship between an account and a post.
type Bookmark struct {
	AccountID uuid.UUID
	PostID    uuid.UUID
	CreatedAt time.Time
}

// Bookmarked represents a post that an account marks.
type Bookmarked struct {
	*Post
}

// NewBookmark creates a new Bookmark instance.
func NewBookmark(accountID, postID string) (*Bookmark, error) {
	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	return &Bookmark{
		AccountID: accountUUID,
		PostID:    postUUID,
		CreatedAt: time.Now(),
	}, nil
}

// createBookmark is a helper function to create new Bookmark instance from the context.
func createBookmark(ctx context.Context, postID string) (*Bookmark, error) {
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}

	accountID := tokenPayload.AccountID.String()
	bookmark, err := NewBookmark(accountID, postID)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}
