-- name: InsertTopic :exec
INSERT INTO topics (
    id, name, slug, description, created_at, updated_at
) VALUES ( 
  $1, $2, $3, $4, $5, $6
);

-- name: SelectTopicByID :one
SELECT * FROM topics
WHERE id = $1
LIMIT 1;

-- name: SelectAllTopics :many
SELECT * FROM topics
ORDER BY name;

-- name: CountTopics :one
SELECT COUNT(id) FROM topics;

-- name: SelectAllTopicsPaginated :many
SELECT * FROM topics
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateTopic :exec
UPDATE topics
SET
  name = COALESCE(sqlc.narg(name), name),
  slug = COALESCE(sqlc.narg(slug), slug),
  description = COALESCE(sqlc.narg(description), description),
  updated_at = sqlc.arg(updated_at)
WHERE id = $1;

-- name: DeleteTopic :execrows
DELETE FROM topics
WHERE id = $1;
