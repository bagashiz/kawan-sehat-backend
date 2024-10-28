// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const deletePost = `-- name: DeletePost :execrows
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deletePost, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const insertPost = `-- name: InsertPost :exec
INSERT INTO posts (
    id, account_id, topic_id, title, content, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
`

type InsertPostParams struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	TopicID   uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) InsertPost(ctx context.Context, arg InsertPostParams) error {
	_, err := q.db.Exec(ctx, insertPost,
		arg.ID,
		arg.AccountID,
		arg.TopicID,
		arg.Title,
		arg.Content,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const selectAllPosts = `-- name: SelectAllPosts :many
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
ORDER BY created_at DESC
`

func (q *Queries) SelectAllPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.Query(ctx, selectAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const selectAllPostsPaginated = `-- name: SelectAllPostsPaginated :many
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`

type SelectAllPostsPaginatedParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) SelectAllPostsPaginated(ctx context.Context, arg SelectAllPostsPaginatedParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, selectAllPostsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const selectPostByID = `-- name: SelectPostByID :one
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
WHERE id = $1
LIMIT 1
`

func (q *Queries) SelectPostByID(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, selectPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.TopicID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectPostsByAccountID = `-- name: SelectPostsByAccountID :many
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
WHERE account_id = $1
ORDER BY created_at DESC
`

func (q *Queries) SelectPostsByAccountID(ctx context.Context, accountID uuid.UUID) ([]Post, error) {
	rows, err := q.db.Query(ctx, selectPostsByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const selectPostsByAccountIDPaginated = `-- name: SelectPostsByAccountIDPaginated :many
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type SelectPostsByAccountIDPaginatedParams struct {
	AccountID uuid.UUID
	Limit     int32
	Offset    int32
}

func (q *Queries) SelectPostsByAccountIDPaginated(ctx context.Context, arg SelectPostsByAccountIDPaginatedParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, selectPostsByAccountIDPaginated, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const selectPostsByTopicID = `-- name: SelectPostsByTopicID :many
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
WHERE topic_id = $1
ORDER BY created_at DESC
`

func (q *Queries) SelectPostsByTopicID(ctx context.Context, topicID uuid.UUID) ([]Post, error) {
	rows, err := q.db.Query(ctx, selectPostsByTopicID, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const selectPostsByTopicIDPaginated = `-- name: SelectPostsByTopicIDPaginated :many
SELECT id, account_id, topic_id, title, content, created_at, updated_at FROM posts
WHERE topic_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type SelectPostsByTopicIDPaginatedParams struct {
	TopicID uuid.UUID
	Limit   int32
	Offset  int32
}

func (q *Queries) SelectPostsByTopicIDPaginated(ctx context.Context, arg SelectPostsByTopicIDPaginatedParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, selectPostsByTopicIDPaginated, arg.TopicID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updatePost = `-- name: UpdatePost :exec
UPDATE posts
SET
  title = COALESCE($2, title),
  content = COALESCE($3, content),
  updated_at = $4
WHERE id = $1
`

type UpdatePostParams struct {
	ID        uuid.UUID
	Title     pgtype.Text
	Content   pgtype.Text
	UpdatedAt time.Time
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.Exec(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.UpdatedAt,
	)
	return err
}
