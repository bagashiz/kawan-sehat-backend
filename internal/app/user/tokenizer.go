package user

import (
	"errors"

	"github.com/google/uuid"
)

// Tokenizer is the interface for interacting with token providers.
type Tokenizer interface {
	CreateToken(payload *TokenPayload) (string, error)
	VerifyToken(token string) (*TokenPayload, error)
}

// TokenPayload contains the information that is stored in the token.
type TokenPayload struct {
	UserRole Role
	ID       uuid.UUID
	UserID   uuid.UUID
}

// newTokenPayload creates a new token payload.
func newTokenPayload(userID uuid.UUID, userRole Role) (*TokenPayload, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate token id")
	}
	return &TokenPayload{
		ID:       id,
		UserID:   userID,
		UserRole: userRole,
	}, nil
}
