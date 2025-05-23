package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	// authHeaderKey is the key used to extract the token from the request header.
	AuthHeaderKey = "Authorization"
	// authType is the type of authentication used in the header.
	AuthType = "Bearer"
	// authPayloadKey is the key used to store the token payload in the request context.
	AuthPayloadKey = "token-payload"
)

// Tokenizer is the interface for interacting with token providers.
type Tokenizer interface {
	CreateToken(payload *TokenPayload, duration time.Duration) (string, time.Time, error)
	VerifyToken(token string) (*TokenPayload, error)
}

// TokenPayload contains the information that is stored in the token.
type TokenPayload struct {
	ID          uuid.UUID
	AccountID   uuid.UUID
	AccountRole Role
}

// newTokenPayload creates a new token payload.
func newTokenPayload(userID uuid.UUID, userRole Role) (*TokenPayload, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate token id")
	}
	return &TokenPayload{
		ID:          id,
		AccountID:   userID,
		AccountRole: userRole,
	}, nil
}

// GetTokenPayload extracts the token payload from the context.
func GetTokenPayload(ctx context.Context) (*TokenPayload, error) {
	payload, ok := ctx.Value(AuthPayloadKey).(*TokenPayload)
	if !ok {
		return nil, ErrAccountUnauthorized
	}
	return payload, nil
}
