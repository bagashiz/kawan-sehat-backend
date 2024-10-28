-- name: InsertBookmark :exec
INSERT INTO bookmarks (
    account_id, post_id, created_at
) VALUES (
  $1, $2, $3
);

-- name: CountAccountBookmarks :one
SELECT COUNT(post_id) FROM bookmarks
WHERE account_id = $1;

-- name: SelectBookmarksByAccountID :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM bookmarks b
JOIN posts p ON b.post_id = p.id
JOIN accounts a ON b.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE b.account_id = $1;

-- name: SelectBookmarksByAccountIDPaginated :many
SELECT p.*, a.username AS account_username, t.name AS topic_name
FROM bookmarks b
JOIN posts p ON b.post_id = p.id
JOIN accounts a ON b.account_id = a.id
JOIN topics t ON p.topic_id = t.id
WHERE b.account_id = $1
LIMIT $2 OFFSET $3;

-- name: DeleteBookmark :execrows
DELETE FROM bookmarks
WHERE account_id = $1 AND post_id = $2;
