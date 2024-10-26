package user

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAccountByID(ctx, uuid)
}
