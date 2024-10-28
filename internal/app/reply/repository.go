package reply

import (
	"context"

	"github.com/google/uuid"
)

// Reader is the interface that provides methods to read reply data from the storage.
type Reader interface {
	GetReplyByID(ctx context.Context, id uuid.UUID) (*Reply, error)
	ListRepliesByCommentID(ctx context.Context, currentID, postID uuid.UUID, limit, page int32) ([]*Reply, int64, error)
	GetVoteReply(ctx context.Context, currentID, postID uuid.UUID) (int16, error)
}

// Writer is the interface that provides methods to write reply data to the storage.
type Writer interface {
	AddReply(ctx context.Context, comment *Reply) error
	DeleteReply(ctx context.Context, id uuid.UUID) error
	VoteReply(ctx context.Context, accountID, replyID uuid.UUID, value int16) (int64, error)
	UpdateVoteReply(ctx context.Context, accountID, replyID uuid.UUID, value int16) (int64, error)
}

// ReadWriter is the interface that combines Reader and Writer interfaces for reply data.
type ReadWriter interface {
	Reader
	Writer
}
