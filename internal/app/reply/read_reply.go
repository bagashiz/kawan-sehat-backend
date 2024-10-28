package reply

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// ListRepliesByCommentID returns a list of comments by post ID.
func (s *Service) ListRepliesByCommentID(
	ctx context.Context, postID string, limit, page int32,
) ([]*Reply, int64, error) {
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return nil, 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, 0, err
	}
	currentID := tokenPayload.AccountID
	return s.repo.ListRepliesByCommentID(ctx, currentID, postUUID, limit, page)
}
