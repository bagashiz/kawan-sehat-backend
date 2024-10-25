package user

import (
	"context"
)

// RegisterAccountParams holds the parameters for the RegisterAccount method.
type RegisterAccountParams struct {
	Username string
	Email    string
	Password string
}

// RegisterAccount creates a new account and stores it in the repository.
func (s *Service) RegisterAccount(ctx context.Context, params RegisterAccountParams) (*Account, error) {
	account, err := New(params.Username, params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	if err := s.repo.AddAccount(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}

// LoginAccountParams holds the parameters for the LoginAccount method.
type LoginAccountParams struct {
	Username string
	Password string
}

// LoginAccount authenticates an account with the given email and password.
// It returns the account info and auth token if the authentication is successful.
func (s *Service) LoginAccount(ctx context.Context, params LoginAccountParams) (*Account, string, error) {
	account, err := s.repo.GetAccountByUsername(ctx, params.Username)
	if err != nil {
		return nil, "", err
	}

	if !account.ComparePassword(params.Password) {
		return nil, "", ErrAccountUnauthorized
	}

	tokenPayload, err := newTokenPayload(account.ID, account.Role)
	if err != nil {
		return nil, "", err
	}

	token, err := s.tokenizer.CreateToken(tokenPayload)
	if err != nil {
		return nil, "", err
	}
	return account, token, nil
}
