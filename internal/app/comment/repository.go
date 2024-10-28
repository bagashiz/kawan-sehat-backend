package comment

import (
	"context"

	"github.com/google/uuid"
)

// Reader is the interface that provides methods to read comment data from the storage.
type Reader interface {
	GetCommentByID(ctx context.Context, id uuid.UUID) (*Comment, error)
	ListCommentsByPostID(ctx context.Context, currentID, postID uuid.UUID, limit, page int32) ([]*Comment, int64, error)
}

// Writer is the interface that provides methods to write comment data to the storage.
type Writer interface {
	AddComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, id uuid.UUID) error
}

// ReadWriter is the interface that combines Reader and Writer interfaces for post data.
type ReadWriter interface {
	Reader
	Writer
}
