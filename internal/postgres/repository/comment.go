package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// AddComment inserts comment data to postgres database.
func (r *PostgresRepository) AddComment(ctx context.Context, c *comment.Comment) error {
	arg := postgres.InsertCommentParams{
		ID:        c.ID,
		PostID:    c.PostID,
		AccountID: c.Account.ID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
	if err := r.db.InsertComment(ctx, arg); err != nil {
		return handleCommentError(err)
	}
	return nil
}

// DeleteComment removes comment data from postgres database.
func (r *PostgresRepository) DeleteComment(ctx context.Context, id uuid.UUID) error {
	count, err := r.db.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	if count == 0 {
		return comment.ErrCommentNotFound
	}
	return nil
}

// GetCommentByID retrieves comment data from postgres database by ID.
func (r *PostgresRepository) GetCommentByID(ctx context.Context, id uuid.UUID) (*comment.Comment, error) {
	result, err := r.db.SelectCommentByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, comment.ErrCommentNotFound
		}
		return nil, err
	}
	comments := result.ToDomain()
	return comments, nil
}

// ListCommentsByPostID returns a list of comments by post ID.
func (r *PostgresRepository) ListCommentsByPostID(
	ctx context.Context,
	currentID, postID uuid.UUID,
	limit, page int32,
) ([]*comment.Comment, int64, error) {
	offset := calculateOffset(limit, page)
	results, err := r.fetchCommentsByPostID(ctx, currentID, postID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.db.CountCommentsByPostID(ctx, postID)
	if err != nil {
		return nil, 0, err
	}
	comments, err := toCommentDomain(results)
	if err != nil {
		return nil, 0, err
	}
	return comments, count, err
}

// fethCommentsByPostID retrieves all comments data from postgres database by post ID.
func (r *PostgresRepository) fetchCommentsByPostID(
	ctx context.Context, currentID, postID uuid.UUID, limit, offset int32,
) (any, error) {
	if limit == 0 {
		return r.db.SelectCommentsByPostID(ctx, postgres.SelectCommentsByPostIDParams{
			AccountID: currentID,
			PostID:    postID,
		})
	}
	return r.db.SelectCommentsByPostIDPaginated(ctx, postgres.SelectCommentsByPostIDPaginatedParams{
		AccountID: currentID,
		PostID:    postID,
		Limit:     limit,
		Offset:    offset,
	})
}

// handleCommentError handles comment postgres repository errors and returns domain errors.
func handleCommentError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return comment.ErrCommentInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			switch pgErr.ConstraintName {
			case "comments_post_id_fkey":
				return post.ErrPostNotFound
			default:
				return err
			}
		default:
			return err
		}
	}
	return err
}
