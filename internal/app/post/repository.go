package post

import (
	"context"

	"github.com/google/uuid"
)

// Reader is the interface that provides methods to read post data from the storage.
type Reader interface {
	GetPostByID(ctx context.Context, accountID, postID uuid.UUID) (*Post, error)
	ListPosts(ctx context.Context, accountID uuid.UUID, limit, page int32) ([]*Post, int64, error)
	ListPostsByTopicID(ctx context.Context, accountID, topicID uuid.UUID, limit, page int32) ([]*Post, int64, error)
	ListPostsByAccountID(ctx context.Context, accountID uuid.UUID, limit, page int32) ([]*Post, int64, error)
	ListAccountBookmarks(ctx context.Context, accountID uuid.UUID, limit, page int32) ([]*Post, int64, error)
	GetVotePost(ctx context.Context, accountID, postID uuid.UUID) (int16, error)
}

// Writer is the interface that provides methods to write post data to the storage.
type Writer interface {
	AddPost(ctx context.Context, post *Post) error
	UpdatePost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, id uuid.UUID) error
	BookmarkPost(ctx context.Context, bookmark *Bookmark) error
	UnbookmarkPost(ctx context.Context, bookmark *Bookmark) error
	VotePost(ctx context.Context, accountID, postID uuid.UUID, value int16) (int64, error)
	UpdateVotePost(ctx context.Context, accountID, postID uuid.UUID, value int16) (int64, error)
}

// ReadWriter is the interface that combines Reader and Writer interfaces for post data.
type ReadWriter interface {
	Reader
	Writer
}
