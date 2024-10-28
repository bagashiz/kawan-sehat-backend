-- name: InsertVotePost :exec
INSERT INTO votes (
    account_id, post_id, value
) VALUES (
  $1, $2, $3
);

-- name: InsertVoteComment :exec
INSERT INTO votes (
    account_id, comment_id, value
) VALUES (
  $1, $2, $3
);

-- name: InsertVoteReply :exec
INSERT INTO votes (
    account_id, reply_id, value
) VALUES (
  $1, $2, $3
);

-- name: SelectVoteByPostID :one
SELECT value FROM votes
WHERE account_id = $1 AND post_id = $2;

-- name: SelectVoteByCommentID :one
SELECT value FROM votes
WHERE account_id = $1 AND comment_id = $2;

-- name: SelectVoteByReplyID :one
SELECT value FROM votes
WHERE account_id = $1 AND reply_id = $2;

-- name: SelectSumVotesByPostID :one
SELECT COALESCE(SUM(value), 0) FROM votes
WHERE post_id = $1;

-- name: UpdateVotePost :exec
UPDATE votes
SET value = $3
WHERE account_id = $1 AND post_id = $2;

-- name: UpdateVoteComment :exec
UPDATE votes
SET value = $3
WHERE account_id = $1 AND comment_id = $2;

-- name: UpdateVoteReply :exec
UPDATE votes
SET value = $3
WHERE account_id = $1 AND reply_id = $2;

-- name: SelectSumVotesByCommentID :one
SELECT COALESCE(SUM(value), 0) FROM votes
WHERE comment_id = $1;

-- name: SelectSumVotesByReplyID :one
SELECT COALESCE(SUM(value), 0) FROM votes
WHERE reply_id = $1;
