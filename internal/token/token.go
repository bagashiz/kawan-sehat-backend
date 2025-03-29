package token

import (
	"fmt"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

// New creates a new token instance based on the provided configuration.
func New(tokenType, secretKey string) (user.Tokenizer, error) {
	switch tokenType {
	case "paseto":
		return NewPasetoTokenMaker(), nil
	case "jwt":
		return NewJWTTokenMaker(secretKey), nil
		// implement more token implementations here
	default:
		return nil, fmt.Errorf("unsupported token type: %s", tokenType)
	}
}
