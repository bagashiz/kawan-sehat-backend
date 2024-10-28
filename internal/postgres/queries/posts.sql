-- name: InsertPost :exec
INSERT INTO posts (
    id, account_id, topic_id, title, content, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
);

-- name: SelectPostByID :one
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.id = $1
LIMIT 1;

-- name: CountPosts :one
SELECT COUNT(id) FROM posts;

-- name: SelectAllPosts :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
ORDER BY p.created_at DESC;

-- name: SelectAllPostsPaginated :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
ORDER BY p.created_at DESC
LIMIT $1
OFFSET $2;

-- name: SelectPostsByAccountID :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.account_id = $1
ORDER BY p.created_at DESC;

-- name: SelectPostsByAccountIDPaginated :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.account_id = $1
ORDER BY p.created_at DESC
LIMIT $2
OFFSET $3;

-- name: SelectPostsByTopicID :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.topic_id = $1
ORDER BY p.created_at DESC;

-- name: SelectPostsByTopicIDPaginated :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM posts p
JOIN accounts a ON p.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE p.topic_id = $1
ORDER BY p.created_at DESC
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
