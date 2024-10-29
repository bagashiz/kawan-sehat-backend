package user

import (
	"context"

	"github.com/google/uuid"
)

// GetAccountByID gets the account by the given ID.
func (s *Service) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAccountByID(ctx, uuid)
}

// ListAccountIllnessHistories lists the account's illness histories.
func (s *Service) ListAccountIllnessHistories(ctx context.Context) ([]*IllnessHistory, error) {
	tokenPayload, err := GetTokenPayload(ctx)
	if err != nil {
		return nil, err
	}
	accountID := tokenPayload.AccountID
	return s.repo.ListIllnessHistoriesByAccountID(ctx, accountID)
}
