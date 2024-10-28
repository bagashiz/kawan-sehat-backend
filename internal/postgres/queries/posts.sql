-- name: InsertPost :exec
INSERT INTO posts (
    id, account_id, topic_id, title, content, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
);

-- name: SelectPostByID :one
SELECT * FROM posts
WHERE id = $1
LIMIT 1;

-- name: CountPosts :one
SELECT COUNT(id) FROM posts;

-- name: SelectAllPosts :many
SELECT * FROM posts
ORDER BY created_at DESC;

-- name: SelectAllPostsPaginated :many
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: CountPostsByAccountID :one
SELECT COUNT(id) FROM posts
WHERE account_id = $1;

-- name: SelectPostsByAccountID :many
SELECT * FROM posts
WHERE account_id = $1
ORDER BY created_at DESC;

-- name: SelectPostsByAccountIDPaginated :many
SELECT * FROM posts
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: CountPostsByTopicID :one
SELECT COUNT(id) FROM posts
WHERE topic_id = $1;

-- name: SelectPostsByTopicID :many
SELECT * FROM posts
WHERE topic_id = $1
ORDER BY created_at DESC;

-- name: SelectPostsByTopicIDPaginated :many
SELECT * FROM posts
WHERE topic_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdatePost :exec
UPDATE posts
SET
  title = COALESCE(sqlc.narg(title), title),
  content = COALESCE(sqlc.narg(content), content),
  updated_at = sqlc.arg(updated_at)
WHERE id = $1;

-- name: DeletePost :execrows
DELETE FROM posts
WHERE id = $1;
