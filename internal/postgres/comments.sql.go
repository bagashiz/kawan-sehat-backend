// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: comments.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const countCommentsByPostID = `-- name: CountCommentsByPostID :one
SELECT COUNT(*) FROM comments
WHERE post_id = $1
`

func (q *Queries) CountCommentsByPostID(ctx context.Context, postID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countCommentsByPostID, postID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteComment = `-- name: DeleteComment :execrows
DELETE FROM comments
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteComment, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const insertComment = `-- name: InsertComment :exec
INSERT INTO comments (
    id, account_id, post_id, content, created_at
) VALUES (
  $1, $2, $3, $4, $5
)
`

type InsertCommentParams struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	PostID    uuid.UUID
	Content   string
	CreatedAt time.Time
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) error {
	_, err := q.db.Exec(ctx, insertComment,
		arg.ID,
		arg.AccountID,
		arg.PostID,
		arg.Content,
		arg.CreatedAt,
	)
	return err
}

const selectCommentByID = `-- name: SelectCommentByID :one
SELECT id, post_id, account_id, content, created_at FROM comments
WHERE id = $1
`

func (q *Queries) SelectCommentByID(ctx context.Context, id uuid.UUID) (Comment, error) {
	row := q.db.QueryRow(ctx, selectCommentByID, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.AccountID,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const selectCommentsByPostID = `-- name: SelectCommentsByPostID :many
SELECT c.id, c.post_id, c.account_id, c.content, c.created_at, 
       a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
       (SELECT COUNT(*) FROM replies r WHERE r.comment_id = c.id) AS total_replies,
       (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.comment_id = c.id) AS total_votes,
       COALESCE((SELECT v.value FROM votes v WHERE v.comment_id = c.id AND v.account_id = $1), 0) AS vote_state
FROM comments c
JOIN accounts a ON c.account_id = a.id
WHERE c.post_id = $2
ORDER BY total_votes
`

type SelectCommentsByPostIDParams struct {
	AccountID uuid.UUID
	PostID    uuid.UUID
}

type SelectCommentsByPostIDRow struct {
	ID              uuid.UUID
	PostID          uuid.UUID
	AccountID       uuid.UUID
	Content         string
	CreatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TotalReplies    int64
	TotalVotes      interface{}
	VoteState       interface{}
}

func (q *Queries) SelectCommentsByPostID(ctx context.Context, arg SelectCommentsByPostIDParams) ([]SelectCommentsByPostIDRow, error) {
	rows, err := q.db.Query(ctx, selectCommentsByPostID, arg.AccountID, arg.PostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectCommentsByPostIDRow{}
	for rows.Next() {
		var i SelectCommentsByPostIDRow
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.AccountID,
			&i.Content,
			&i.CreatedAt,
			&i.AccountUsername,
			&i.AccountAvatar,
			&i.AccountRole,
			&i.TotalReplies,
			&i.TotalVotes,
			&i.VoteState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectCommentsByPostIDPaginated = `-- name: SelectCommentsByPostIDPaginated :many
SELECT c.id, c.post_id, c.account_id, c.content, c.created_at,
    a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
    (SELECT COUNT(*) FROM replies r WHERE r.comment_id = c.id) AS total_replies,
    (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.comment_id = c.id) AS total_votes,
       COALESCE((SELECT v.value FROM votes v WHERE v.comment_id = c.id AND v.account_id = $1), 0) AS vote_state
FROM comments c
JOIN accounts a ON c.account_id = a.id
WHERE c.post_id = $2
ORDER BY total_votes
LIMIT $3
OFFSET $4
`

type SelectCommentsByPostIDPaginatedParams struct {
	AccountID uuid.UUID
	PostID    uuid.UUID
	Limit     int32
	Offset    int32
}

type SelectCommentsByPostIDPaginatedRow struct {
	ID              uuid.UUID
	PostID          uuid.UUID
	AccountID       uuid.UUID
	Content         string
	CreatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TotalReplies    int64
	TotalVotes      interface{}
	VoteState       interface{}
}

func (q *Queries) SelectCommentsByPostIDPaginated(ctx context.Context, arg SelectCommentsByPostIDPaginatedParams) ([]SelectCommentsByPostIDPaginatedRow, error) {
	rows, err := q.db.Query(ctx, selectCommentsByPostIDPaginated,
		arg.AccountID,
		arg.PostID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectCommentsByPostIDPaginatedRow{}
	for rows.Next() {
		var i SelectCommentsByPostIDPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.AccountID,
			&i.Content,
			&i.CreatedAt,
			&i.AccountUsername,
			&i.AccountAvatar,
			&i.AccountRole,
			&i.TotalReplies,
			&i.TotalVotes,
			&i.VoteState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
