// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const insertAccount = `-- name: InsertAccount :exec
INSERT INTO accounts (
    id, full_name, username,
    nik, email, password,
    gender, role, avatar,
    created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6,
  $7, $8, $9, $10, $11
)
`

type InsertAccountParams struct {
	ID        uuid.UUID
	FullName  pgtype.Text
	Username  string
	Nik       pgtype.Text
	Email     string
	Password  string
	Gender    AccountGender
	Role      AccountRole
	Avatar    AccountAvatar
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) InsertAccount(ctx context.Context, arg InsertAccountParams) error {
	_, err := q.db.Exec(ctx, insertAccount,
		arg.ID,
		arg.FullName,
		arg.Username,
		arg.Nik,
		arg.Email,
		arg.Password,
		arg.Gender,
		arg.Role,
		arg.Avatar,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const selectAccountByID = `-- name: SelectAccountByID :one
SELECT id, full_name, nik, username, email, password, gender, role, avatar, created_at, updated_at FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) SelectAccountByID(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRow(ctx, selectAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Nik,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Gender,
		&i.Role,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectAccountByUsername = `-- name: SelectAccountByUsername :one
SELECT id, full_name, nik, username, email, password, gender, role, avatar, created_at, updated_at FROM accounts
WHERE username = $1
LIMIT 1
`

func (q *Queries) SelectAccountByUsername(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRow(ctx, selectAccountByUsername, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Nik,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Gender,
		&i.Role,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectIllnessHistoriesByAccountID = `-- name: SelectIllnessHistoriesByAccountID :many
SELECT account_id, illness, date FROM illness_histories
WHERE account_id = $1
`

func (q *Queries) SelectIllnessHistoriesByAccountID(ctx context.Context, accountID uuid.UUID) ([]IllnessHistory, error) {
	rows, err := q.db.Query(ctx, selectIllnessHistoriesByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []IllnessHistory{}
	for rows.Next() {
		var i IllnessHistory
		if err := rows.Scan(&i.AccountID, &i.Illness, &i.Date); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE accounts
SET
  full_name = COALESCE($2, full_name),
  username = COALESCE($3, username),
  nik = COALESCE($4, nik),
  email = COALESCE($5, email),
  password = COALESCE($6, password),
  gender = COALESCE($7, gender),
  role = COALESCE($8, role),
  avatar = COALESCE($9, avatar),
  updated_at = $10
WHERE id = $1
`

type UpdateAccountParams struct {
	ID        uuid.UUID
	FullName  pgtype.Text
	Username  pgtype.Text
	Nik       pgtype.Text
	Email     pgtype.Text
	Password  pgtype.Text
	Gender    NullAccountGender
	Role      NullAccountRole
	Avatar    NullAccountAvatar
	UpdatedAt time.Time
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.Exec(ctx, updateAccount,
		arg.ID,
		arg.FullName,
		arg.Username,
		arg.Nik,
		arg.Email,
		arg.Password,
		arg.Gender,
		arg.Role,
		arg.Avatar,
		arg.UpdatedAt,
	)
	return err
}
