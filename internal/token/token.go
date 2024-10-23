package token

import (
	"fmt"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/config"
)

// New creates a new token instance based on
// the provided configuration.
func New(cfg *config.Token) (user.Tokenizer, error) {
	switch cfg.Type {
	case "paseto":
		return NewPaseto(cfg)
		// implement more token implementations here
	default:
		return nil, fmt.Errorf("unsupported token type: %s", cfg.Type)
	}
}
