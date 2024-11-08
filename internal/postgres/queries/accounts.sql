-- name: InsertAccount :exec
INSERT INTO accounts (
    id, full_name, username,
    nik, email, password,
    gender, role, avatar,
    created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6,
  $7, $8, $9, $10, $11
);

-- name: UpdateAccount :exec
UPDATE accounts
SET
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  username = COALESCE(sqlc.narg(username), username),
  nik = COALESCE(sqlc.narg(nik), nik),
  email = COALESCE(sqlc.narg(email), email),
  password = COALESCE(sqlc.narg(password), password),
  gender = COALESCE(sqlc.narg(gender), gender),
  role = COALESCE(sqlc.narg(role), role),
  avatar = COALESCE(sqlc.narg(avatar), avatar),
  updated_at = sqlc.arg(updated_at)
WHERE id = $1;

-- name: SelectAccountByUsername :one
SELECT * FROM accounts
WHERE username = $1
LIMIT 1;

-- name: SelectAccountByID :one
SELECT * FROM accounts
WHERE id = $1
LIMIT 1;

-- name: SelectIllnessHistoriesByAccountID :many
SELECT * FROM illness_histories
WHERE account_id = $1
ORDER BY date DESC;
