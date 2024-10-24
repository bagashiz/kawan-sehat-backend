package user

import (
	"context"
)

// RegisterAccountParams holds the parameters for the RegisterAccount method.
type RegisterAccountParams struct {
	Username string `validate:"required,username,lte=30"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

// RegisterAccount creates a new account and stores it in the repository.
func (s *Service) RegisterAccount(ctx context.Context, params RegisterAccountParams) (*Account, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	account, err := New(params.Username, params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	if err := s.repo.AddAccount(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}
