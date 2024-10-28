package comment

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// ListCommentsByPostID returns a list of comments by post ID.
func (s *Service) ListCommentsByPostID(
	ctx context.Context, postID string, limit, page int32,
) ([]*Comment, int64, error) {
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return nil, 0, err
	}
	tokenPayload, err := user.GetTokenPayload(ctx)
	if err != nil {
		return nil, 0, err
	}
	currentID := tokenPayload.AccountID
	return s.repo.ListCommentsByPostID(ctx, currentID, postUUID, limit, page)
}
