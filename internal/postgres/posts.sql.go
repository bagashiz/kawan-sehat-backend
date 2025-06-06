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

const countPosts = `-- name: CountPosts :one
SELECT COUNT(id) FROM posts
`

func (q *Queries) CountPosts(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countPosts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByAccountID = `-- name: CountPostsByAccountID :one
SELECT COUNT(id) FROM posts
WHERE account_id = $1
`

func (q *Queries) CountPostsByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countPostsByAccountID, accountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByTopicID = `-- name: CountPostsByTopicID :one
SELECT COUNT(id) FROM posts
WHERE topic_id = $1
`

func (q *Queries) CountPostsByTopicID(ctx context.Context, topicID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countPostsByTopicID, topicID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

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
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
ORDER BY p.created_at DESC
`

type SelectAllPostsRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectAllPosts(ctx context.Context, accountID uuid.UUID) ([]SelectAllPostsRow, error) {
	rows, err := q.db.Query(ctx, selectAllPosts, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectAllPostsRow{}
	for rows.Next() {
		var i SelectAllPostsRow
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
			&i.AccountRole,
			&i.TopicName,
			&i.TotalComments,
			&i.TotalVotes,
			&i.VoteState,
			&i.IsBookmarked,
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
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
ORDER BY p.created_at DESC
LIMIT $2
OFFSET $3
`

type SelectAllPostsPaginatedParams struct {
	AccountID uuid.UUID
	Limit     int32
	Offset    int32
}

type SelectAllPostsPaginatedRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectAllPostsPaginated(ctx context.Context, arg SelectAllPostsPaginatedParams) ([]SelectAllPostsPaginatedRow, error) {
	rows, err := q.db.Query(ctx, selectAllPostsPaginated, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectAllPostsPaginatedRow{}
	for rows.Next() {
		var i SelectAllPostsPaginatedRow
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
			&i.AccountRole,
			&i.TopicName,
			&i.TotalComments,
			&i.TotalVotes,
			&i.VoteState,
			&i.IsBookmarked,
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
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.id = $2
LIMIT 1
`

type SelectPostByIDParams struct {
	AccountID uuid.UUID
	ID        uuid.UUID
}

type SelectPostByIDRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectPostByID(ctx context.Context, arg SelectPostByIDParams) (SelectPostByIDRow, error) {
	row := q.db.QueryRow(ctx, selectPostByID, arg.AccountID, arg.ID)
	var i SelectPostByIDRow
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.TopicID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountUsername,
		&i.AccountAvatar,
		&i.AccountRole,
		&i.TopicName,
		&i.TotalComments,
		&i.TotalVotes,
		&i.VoteState,
		&i.IsBookmarked,
	)
	return i, err
}

const selectPostsByAccountID = `-- name: SelectPostsByAccountID :many
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.account_id = $1
ORDER BY p.created_at DESC
`

type SelectPostsByAccountIDRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectPostsByAccountID(ctx context.Context, accountID uuid.UUID) ([]SelectPostsByAccountIDRow, error) {
	rows, err := q.db.Query(ctx, selectPostsByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectPostsByAccountIDRow{}
	for rows.Next() {
		var i SelectPostsByAccountIDRow
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
			&i.AccountRole,
			&i.TopicName,
			&i.TotalComments,
			&i.TotalVotes,
			&i.VoteState,
			&i.IsBookmarked,
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
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.account_id = $1
ORDER BY p.created_at DESC
LIMIT $2
OFFSET $3
`

type SelectPostsByAccountIDPaginatedParams struct {
	AccountID uuid.UUID
	Limit     int32
	Offset    int32
}

type SelectPostsByAccountIDPaginatedRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectPostsByAccountIDPaginated(ctx context.Context, arg SelectPostsByAccountIDPaginatedParams) ([]SelectPostsByAccountIDPaginatedRow, error) {
	rows, err := q.db.Query(ctx, selectPostsByAccountIDPaginated, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectPostsByAccountIDPaginatedRow{}
	for rows.Next() {
		var i SelectPostsByAccountIDPaginatedRow
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
			&i.AccountRole,
			&i.TopicName,
			&i.TotalComments,
			&i.TotalVotes,
			&i.VoteState,
			&i.IsBookmarked,
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
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.topic_id = $2
ORDER BY p.created_at DESC
`

type SelectPostsByTopicIDParams struct {
	AccountID uuid.UUID
	TopicID   uuid.UUID
}

type SelectPostsByTopicIDRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectPostsByTopicID(ctx context.Context, arg SelectPostsByTopicIDParams) ([]SelectPostsByTopicIDRow, error) {
	rows, err := q.db.Query(ctx, selectPostsByTopicID, arg.AccountID, arg.TopicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectPostsByTopicIDRow{}
	for rows.Next() {
		var i SelectPostsByTopicIDRow
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
			&i.AccountRole,
			&i.TopicName,
			&i.TotalComments,
			&i.TotalVotes,
			&i.VoteState,
			&i.IsBookmarked,
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
SELECT p.id, p.account_id, p.topic_id, p.title, p.content, p.created_at, p.updated_at,
  a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
  t.name AS topic_name,
  (SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS total_comments,
  (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.post_id = p.id) AS total_votes,
  COALESCE((SELECT v.value FROM votes v WHERE v.post_id = p.id AND v.account_id = $1), 0) AS vote_state,
  CASE WHEN EXISTS (SELECT 1 FROM bookmarks b WHERE b.post_id = p.id AND b.account_id = $1) THEN TRUE ELSE FALSE END AS is_bookmarked
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.topic_id = $2
ORDER BY p.created_at DESC
LIMIT $3
OFFSET $4
`

type SelectPostsByTopicIDPaginatedParams struct {
	AccountID uuid.UUID
	TopicID   uuid.UUID
	Limit     int32
	Offset    int32
}

type SelectPostsByTopicIDPaginatedRow struct {
	ID              uuid.UUID
	AccountID       uuid.UUID
	TopicID         uuid.UUID
	Title           string
	Content         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccountUsername string
	AccountAvatar   AccountAvatar
	AccountRole     AccountRole
	TopicName       string
	TotalComments   int64
	TotalVotes      interface{}
	VoteState       interface{}
	IsBookmarked    bool
}

func (q *Queries) SelectPostsByTopicIDPaginated(ctx context.Context, arg SelectPostsByTopicIDPaginatedParams) ([]SelectPostsByTopicIDPaginatedRow, error) {
	rows, err := q.db.Query(ctx, selectPostsByTopicIDPaginated,
		arg.AccountID,
		arg.TopicID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectPostsByTopicIDPaginatedRow{}
	for rows.Next() {
		var i SelectPostsByTopicIDPaginatedRow
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
			&i.AccountRole,
			&i.TopicName,
			&i.TotalComments,
			&i.TotalVotes,
			&i.VoteState,
			&i.IsBookmarked,
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
