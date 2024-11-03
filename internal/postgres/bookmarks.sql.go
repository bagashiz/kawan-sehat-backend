// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: bookmarks.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const countAccountBookmarks = `-- name: CountAccountBookmarks :one
SELECT COUNT(post_id) FROM bookmarks
WHERE account_id = $1
`

func (q *Queries) CountAccountBookmarks(ctx context.Context, accountID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countAccountBookmarks, accountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteBookmark = `-- name: DeleteBookmark :execrows
DELETE FROM bookmarks
WHERE account_id = $1 AND post_id = $2
`

type DeleteBookmarkParams struct {
	AccountID uuid.UUID
	PostID    uuid.UUID
}

func (q *Queries) DeleteBookmark(ctx context.Context, arg DeleteBookmarkParams) (int64, error) {
	result, err := q.db.Exec(ctx, deleteBookmark, arg.AccountID, arg.PostID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const insertBookmark = `-- name: InsertBookmark :exec
INSERT INTO bookmarks (
    account_id, post_id, created_at
) VALUES (
  $1, $2, $3
)
`

type InsertBookmarkParams struct {
	AccountID uuid.UUID
	PostID    uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) InsertBookmark(ctx context.Context, arg InsertBookmarkParams) error {
	_, err := q.db.Exec(ctx, insertBookmark, arg.AccountID, arg.PostID, arg.CreatedAt)
	return err
}

const selectBookmarksByAccountID = `-- name: SelectBookmarksByAccountID :many
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at, 
  a.username AS account_username, a.avatar AS account_avatar, t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state
FROM bookmarks b
JOIN posts p ON b.post_id = p.id
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE b.account_id = $1
`

type SelectBookmarksByAccountIDRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
}

func (q *Queries) SelectBookmarksByAccountID(ctx context.Context, accountID uuid.UUID) ([]SelectBookmarksByAccountIDRow, error) {
	rows, err := q.db.Query(ctx, selectBookmarksByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectBookmarksByAccountIDRow{}
	for rows.Next() {
		var i SelectBookmarksByAccountIDRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccountUsername,
			&i.AccountAvatar,
			&i.TopicName,
			&i.TotalComments,
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

const selectBookmarksByAccountIDPaginated = `-- name: SelectBookmarksByAccountIDPaginated :many
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at, 
  a.username AS account_username, a.avatar AS account_avatar, t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state
FROM bookmarks b
JOIN posts p ON b.post_id = p.id
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE b.account_id = $1
LIMIT $2 OFFSET $3
`

type SelectBookmarksByAccountIDPaginatedParams struct {
	AccountID uuid.UUID
	Limit     int32
	Offset    int32
}

type SelectBookmarksByAccountIDPaginatedRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
}

func (q *Queries) SelectBookmarksByAccountIDPaginated(ctx context.Context, arg SelectBookmarksByAccountIDPaginatedParams) ([]SelectBookmarksByAccountIDPaginatedRow, error) {
	rows, err := q.db.Query(ctx, selectBookmarksByAccountIDPaginated, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectBookmarksByAccountIDPaginatedRow{}
	for rows.Next() {
		var i SelectBookmarksByAccountIDPaginatedRow
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccountUsername,
			&i.AccountAvatar,
			&i.TopicName,
			&i.TotalComments,
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
