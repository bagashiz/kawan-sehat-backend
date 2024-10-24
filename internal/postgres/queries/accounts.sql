-- name: InsertAccount :exec
INSERT INTO accounts (
    id, full_name, username,
    nik, email, password,
    gender, role, avatar,
    illness_history,
    created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6,
  $7, $8, $9, $10, $11, $12
);

-- name: SelectAccountByUsername :one
SELECT * FROM accounts
WHERE username = $1
LIMIT 1;
