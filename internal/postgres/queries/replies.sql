-- name: InsertReply :exec
INSERT INTO replies (
    id, account_id, comment_id, content, created_at
) VALUES (
  $1, $2, $3, $4, $5
);

-- name: SelectReplyByID :one
SELECT * FROM replies
WHERE id = $1;

-- name: CountRepliesByCommentID :one
SELECT COUNT(*) FROM replies
WHERE comment_id = $1;

-- name: SelectRepliesByCommentID :many
SELECT r.*, 
       a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
       (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.reply_id = r.id) AS total_votes,
       COALESCE((SELECT v.value FROM votes v WHERE v.reply_id = r.id AND v.account_id = $1), 0) AS vote_state
FROM replies r
JOIN accounts a ON r.account_id = a.id
WHERE r.comment_id = $2
ORDER BY total_votes;

-- name: SelectRepliesByCommentIDPaginated :many
SELECT r.*,
       a.username AS account_username, a.avatar AS account_avatar, a.role AS account_role, 
       (SELECT COALESCE(SUM(v.value), 0) FROM votes v WHERE v.reply_id = r.id) AS total_votes,
       COALESCE((SELECT v.value FROM votes v WHERE v.reply_id = r.id AND v.account_id = $1), 0) AS vote_state
FROM replies r
JOIN accounts a ON r.account_id = a.id
WHERE r.comment_id = $2
ORDER BY total_votes
LIMIT $3
OFFSET $4;

-- name: DeleteReply :execrows
DELETE FROM replies
WHERE id = $1;
