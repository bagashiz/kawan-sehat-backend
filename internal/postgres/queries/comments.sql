-- name: InsertComment :exec
INSERT INTO comments (
    id, account_id, post_id, content, created_at
) VALUES (
  $1, $2, $3, $4, $5
);

-- name: SelectCommentByID :one
SELECT * FROM comments
WHERE id = $1;

-- name: CountCommentsByPostID :one
SELECT COUNT(*) FROM comments
WHERE post_id = $1;

-- name: SelectCommentsByPostID :many
SELECT c.*, 
       a.username AS account_username,
       (SELECT COUNT(*) FROM replies r WHERE r.comment_id = c.id) AS total_replies,
       (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.comment_id = c.id) AS total_votes,
       COALESCE((SELECT v.value FROM votes v WHERE v.comment_id = c.id AND v.account_id = $1), 0) AS vote_state
FROM comments c
JOIN accounts a ON c.account_id = a.id
WHERE c.post_id = $2
ORDER BY total_votes;

-- name: SelectCommentsByPostIDPaginated :many
SELECT c.*, a.username AS account_username,
    (SELECT COUNT(*) FROM replies r WHERE r.comment_id = c.id) AS total_replies,
    (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.comment_id = c.id) AS total_votes,
       COALESCE((SELECT v.value FROM votes v WHERE v.comment_id = c.id AND v.account_id = $1), 0) AS vote_state
FROM comments c
JOIN accounts a ON c.account_id = a.id
WHERE c.post_id = $2
ORDER BY total_votes
LIMIT $3
OFFSET $4;

-- name: DeleteComment :execrows
DELETE FROM comments
WHERE id = $1;
