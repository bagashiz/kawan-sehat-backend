package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/reply"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// AddReply inserts comment data to postgres database.
func (r *PostgresRepository) AddReply(ctx context.Context, c *reply.Reply) error {
	arg := postgres.InsertReplyParams{
		ID:        c.ID,
		CommentID: c.CommentID,
		AccountID: c.Account.ID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
	if err := r.db.InsertReply(ctx, arg); err != nil {
		return handleReplyError(err)
	}
	return nil
}

// DeleteReply removes comment data from postgres database.
func (r *PostgresRepository) DeleteReply(ctx context.Context, id uuid.UUID) error {
	count, err := r.db.DeleteReply(ctx, id)
	if err != nil {
		return err
	}
	if count == 0 {
		return reply.ErrReplyNotFound
	}
	return nil
}

// GetReplyByID retrieves comment data from postgres database by ID.
func (r *PostgresRepository) GetReplyByID(ctx context.Context, id uuid.UUID) (*reply.Reply, error) {
	result, err := r.db.SelectReplyByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, reply.ErrReplyNotFound
		}
		return nil, err
	}
	reply := result.ToDomain()
	return reply, nil
}

// ListRepliesByCommentID returns a list of comments by post ID.
func (r *PostgresRepository) ListRepliesByCommentID(
	ctx context.Context,
	currentID, commentID uuid.UUID,
	limit, page int32,
) ([]*reply.Reply, int64, error) {
	offset := calculateOffset(limit, page)
	results, err := r.fetchRepliesByCommentID(ctx, currentID, commentID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.db.CountRepliesByCommentID(ctx, commentID)
	if err != nil {
		return nil, 0, err
	}
	comments, err := toReplyDomain(results)
	if err != nil {
		return nil, 0, err
	}
	return comments, count, err
}

// fethRepliesByReplyID retrieves all comments data from postgres database by post ID.
func (r *PostgresRepository) fetchRepliesByCommentID(
	ctx context.Context, currentID, commentID uuid.UUID, limit, offset int32,
) (any, error) {
	if limit == 0 {
		return r.db.SelectRepliesByCommentID(ctx, postgres.SelectRepliesByCommentIDParams{
			AccountID: currentID,
			CommentID: commentID,
		})
	}
	return r.db.SelectRepliesByCommentIDPaginated(ctx, postgres.SelectRepliesByCommentIDPaginatedParams{
		AccountID: currentID,
		CommentID: commentID,
		Limit:     limit,
		Offset:    offset,
	})
}

// handleReplyError handles comment postgres repository errors and returns domain errors.
func handleReplyError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return reply.ErrReplyInvalid
		default:
			return err
		}
	}
	return err
}
