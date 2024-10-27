package user

import (
	"context"
)

// UpdateAccountParams holds the parameters for the UpdateAccount method.
type UpdateAccountParams struct {
	FullName       string
	Username       string
	NIK            string
	Email          string
	Password       string
	Gender         string
	Role           string
	Avatar         string
	IllnessHistory string
}

// UpdateAccount updates the account with the given parameters.
func (s *Service) UpdateAccount(ctx context.Context, params UpdateAccountParams) (*Account, error) {
	tokenPayload, err := GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	accountID := tokenPayload.AccountID

	current, err := s.repo.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	if err := current.Update(
		params.FullName, params.Username, params.NIK,
		params.Email, params.Password, params.Gender,
		params.Role, params.Avatar, params.IllnessHistory,
	); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateAccount(ctx, current); err != nil {
		return nil, err
	}

	return current, nil
}
