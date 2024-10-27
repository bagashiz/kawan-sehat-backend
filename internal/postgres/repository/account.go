package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// AddAccount inserts user account data to postgres database.
func (r *PostgresRepository) AddAccount(ctx context.Context, account *user.Account) error {
	arg := postgres.InsertAccountParams{
		ID: account.ID,
		FullName: pgtype.Text{
			String: account.FullName, Valid: account.FullName != "",
		},
		Nik: pgtype.Text{
			String: account.NIK, Valid: account.NIK != "" && len(account.NIK) <= 16,
		},
		Username: account.Username,
		Email:    account.Email,
		Password: account.Password,
		Gender:   postgres.AccountGender(account.Gender),
		Role:     postgres.AccountRole(account.Role),
		Avatar:   postgres.AccountAvatar(account.Avatar),
		IllnessHistory: pgtype.Text{
			String: account.IllnessHistory, Valid: account.IllnessHistory != "",
		},
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}

	if err := r.db.InsertAccount(ctx, arg); err != nil {
		return handleAccountError(err)
	}

	return nil
}

// GetAccountByUsername retrieves user account data from postgres database by username.
func (r *PostgresRepository) GetAccountByUsername(ctx context.Context, username string) (*user.Account, error) {
	result, err := r.db.SelectAccountByUsername(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user.ErrAccountNotFound
		}
		return nil, err
	}

	account := result.ToDomain()

	return account, nil
}

// GetAccountByID retrieves user account data from postgres database by ID.
func (r *PostgresRepository) GetAccountByID(ctx context.Context, id uuid.UUID) (*user.Account, error) {
	result, err := r.db.SelectAccountByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user.ErrAccountNotFound
		}
		return nil, err
	}

	account := result.ToDomain()

	return account, nil
}

// handleAccountError handles account postgres repository errors and returns domain errors.
func handleAccountError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return user.ErrAccountInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			switch pgErr.ConstraintName {
			case "accounts_username_key":
				return user.ErrAccountDuplicateUsername
			case "accounts_email_key":
				return user.ErrAccountDuplicateEmail
			default:
				return pgErr
			}
		default:
			return pgErr
		}
	}
	return err
}
