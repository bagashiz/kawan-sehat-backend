package repository

import (
	"context"
	"errors"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/reply"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// VotePost adds a vote to a post in postgres database.
func (r *PostgresRepository) VotePost(
	ctx context.Context, accountID, postID uuid.UUID, value int16,
) (int64, error) {
	arg := postgres.InsertVotePostParams{
		AccountID: accountID,
		PostID:    pgtype.UUID{Bytes: postID, Valid: postID != uuid.Nil},
		Value:     value,
	}
	if err := r.db.InsertVotePost(ctx, arg); err != nil {
		return 0, handleVoteError(err)
	}
	result, err := r.db.SelectSumVotesByPostID(
		ctx, pgtype.UUID{Bytes: postID, Valid: postID != uuid.Nil},
	)
	if err != nil {
		return 0, handleVoteError(err)
	}
	count, ok := result.(int64)
	if !ok {
		return 0, errors.New("something went wrong")
	}
	return count, nil
}

// VoteComment adds a vote to a comment in the postgres database.
func (r *PostgresRepository) VoteComment(
	ctx context.Context, accountID, commentID uuid.UUID, value int16,
) (int64, error) {
	arg := postgres.InsertVoteCommentParams{
		AccountID: accountID,
		CommentID: pgtype.UUID{Bytes: commentID, Valid: commentID != uuid.Nil},
		Value:     value,
	}
	if err := r.db.InsertVoteComment(ctx, arg); err != nil {
		return 0, handleVoteError(err)
	}
	result, err := r.db.SelectSumVotesByCommentID(
		ctx, pgtype.UUID{Bytes: commentID, Valid: commentID != uuid.Nil},
	)
	if err != nil {
		return 0, handleVoteError(err)
	}
	count, ok := result.(int64)
	if !ok {
		return 0, errors.New("something went wrong")
	}
	return count, nil
}

// VoteReply adds a vote to a reply in the postgres database.
func (r *PostgresRepository) VoteReply(
	ctx context.Context, accountID, replyID uuid.UUID, value int16,
) (int64, error) {
	arg := postgres.InsertVoteReplyParams{
		AccountID: accountID,
		ReplyID:   pgtype.UUID{Bytes: replyID, Valid: replyID != uuid.Nil},
		Value:     value,
	}
	if err := r.db.InsertVoteReply(ctx, arg); err != nil {
		return 0, handleVoteError(err)
	}
	result, err := r.db.SelectSumVotesByReplyID(
		ctx, pgtype.UUID{Bytes: replyID, Valid: replyID != uuid.Nil},
	)
	if err != nil {
		return 0, handleVoteError(err)
	}
	count, ok := result.(int64)
	if !ok {
		return 0, errors.New("something went wrong")
	}
	return count, nil
}

// GetVotePost returns a vote from a post in the postgres database.
func (r *PostgresRepository) GetVotePost(
	ctx context.Context, accountID, postID uuid.UUID,
) (int16, error) {
	arg := postgres.SelectVoteByPostIDParams{
		AccountID: accountID,
		PostID:    pgtype.UUID{Bytes: postID, Valid: postID != uuid.Nil},
	}
	value, err := r.db.SelectVoteByPostID(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, post.ErrPostVoteNotFound
		}
		return 0, handleVoteError(err)
	}
	return value, nil
}

// GetVoteComment returns a vote from a comment in the postgres database.
func (r *PostgresRepository) GetVoteComment(
	ctx context.Context, accountID, commentID uuid.UUID,
) (int16, error) {
	arg := postgres.SelectVoteByCommentIDParams{
		AccountID: accountID,
		CommentID: pgtype.UUID{Bytes: commentID, Valid: commentID != uuid.Nil},
	}
	value, err := r.db.SelectVoteByCommentID(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, comment.ErrCommentVoteNotFound
		}
		return 0, handleVoteError(err)
	}
	return value, nil
}

// GetVoteReply returns a vote from a reply in the postgres database.
func (r *PostgresRepository) GetVoteReply(
	ctx context.Context, accountID, replyID uuid.UUID,
) (int16, error) {
	arg := postgres.SelectVoteByReplyIDParams{
		AccountID: accountID,
		ReplyID:   pgtype.UUID{Bytes: replyID, Valid: replyID != uuid.Nil},
	}
	value, err := r.db.SelectVoteByReplyID(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, reply.ErrReplyVoteNotFound
		}
		return 0, handleVoteError(err)
	}
	return value, nil
}

// UpdateVotePost updates a vote from a post in the postgres database.
func (r *PostgresRepository) UpdateVotePost(
	ctx context.Context, accountID, postID uuid.UUID, value int16,
) (int64, error) {
	arg := postgres.UpdateVotePostParams{
		AccountID: accountID,
		PostID:    pgtype.UUID{Bytes: postID, Valid: postID != uuid.Nil},
		Value:     value,
	}
	if err := r.db.UpdateVotePost(ctx, arg); err != nil {
		return 0, handleVoteError(err)
	}
	result, err := r.db.SelectSumVotesByPostID(
		ctx, pgtype.UUID{Bytes: postID, Valid: postID != uuid.Nil},
	)
	if err != nil {
		return 0, handleVoteError(err)
	}
	count, ok := result.(int64)
	if !ok {
		return 0, errors.New("something went wrong")
	}
	return count, nil
}

// UpdateVoteComment updates a vote from a comment in the postgres database.
func (r *PostgresRepository) UpdateVoteComment(
	ctx context.Context, accountID, commentID uuid.UUID, value int16,
) (int64, error) {
	arg := postgres.UpdateVoteCommentParams{
		AccountID: accountID,
		CommentID: pgtype.UUID{Bytes: commentID, Valid: commentID != uuid.Nil},
		Value:     value,
	}
	if err := r.db.UpdateVoteComment(ctx, arg); err != nil {
		return 0, handleVoteError(err)
	}
	result, err := r.db.SelectSumVotesByCommentID(
		ctx, pgtype.UUID{Bytes: commentID, Valid: commentID != uuid.Nil},
	)
	if err != nil {
		return 0, handleVoteError(err)
	}
	count, ok := result.(int64)
	if !ok {
		return 0, errors.New("something went wrong")
	}
	return count, nil
}

// UpdateVoteReply updates a vote from a reply in the postgres database.
func (r *PostgresRepository) UpdateVoteReply(
	ctx context.Context, accountID, replyID uuid.UUID, value int16,
) (int64, error) {
	arg := postgres.UpdateVoteReplyParams{
		AccountID: accountID,
		ReplyID:   pgtype.UUID{Bytes: replyID, Valid: replyID != uuid.Nil},
		Value:     value,
	}
	if err := r.db.UpdateVoteReply(ctx, arg); err != nil {
		return 0, handleVoteError(err)
	}
	result, err := r.db.SelectSumVotesByReplyID(
		ctx, pgtype.UUID{Bytes: replyID, Valid: replyID != uuid.Nil},
	)
	if err != nil {
		return 0, handleVoteError(err)
	}
	count, ok := result.(int64)
	if !ok {
		return 0, errors.New("something went wrong")
	}
	return count, nil
}

// handleVoteError handles post postgres repository errors and returns domain errors.
func handleVoteError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return errors.New("vote data is invalid")
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			switch pgErr.ConstraintName {
			case "votes_account_id_fkey":
				return user.ErrAccountNotFound
			case "votes_post_id_fkey":
				return post.ErrPostNotFound
			case "votes_comment_id_fkey":
				return comment.ErrCommentNotFound
			case "votes_reply_id_fkey":
				return reply.ErrReplyNotFound
			case "votes_account_id_post_id_idx":
				return post.ErrPostVoteAlreadyVoted
			case "votes_account_id_comment_id_idx":
				return comment.ErrCommentVoteAlreadyVoted
			case "votes_account_id_reply_id_idx":
				return reply.ErrReplyVoteAlreadyVoted
			default:
				return err
			}
		}
	}
	return err
}
