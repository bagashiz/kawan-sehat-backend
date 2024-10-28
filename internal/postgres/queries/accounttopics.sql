-- name: InsertAccountTopic :exec
INSERT INTO account_topics (
    account_id, topic_id, created_at
) VALUES ( 
  $1, $2, $3
);

-- name: CountTopicsByAccountID :one
SELECT COUNT(topic_id) FROM account_topics
WHERE account_id = $1;

-- name: SelectTopicsByAccountID :many
SELECT t.* FROM topics t
JOIN account_topics at
ON t.id = at.topic_id
WHERE at.account_id = $1;

-- name: SelectTopicsByAccountIDPaginated :many
SELECT t.* FROM topics t
JOIN account_topics at
ON t.id = at.topic_id
WHERE at.account_id = $1
LIMIT $2 OFFSET $3;

-- name: DeleteAccountTopic :execrows
DELETE FROM account_topics
WHERE account_id = $1 AND topic_id = $2;
