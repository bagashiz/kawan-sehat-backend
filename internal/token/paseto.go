package token

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/config"
)

// Paseto implements user.Tokenizer interface
// and provides an access to the paseto library.
type Paseto struct {
	token    *paseto.Token
	key      *paseto.V4SymmetricKey
	parser   *paseto.Parser
	duration time.Duration
}

// NewPaseto creates a new paseto instance.
func NewPaseto(cfg *config.Token) (*Paseto, error) {
	duration, err := time.ParseDuration(cfg.Duration)
	if err != nil {
		return nil, errors.New("invalid token duration")
	}

	token := paseto.NewToken()
	key := paseto.NewV4SymmetricKey()
	parser := paseto.NewParser()

	return &Paseto{
		&token,
		&key,
		&parser,
		duration,
	}, nil
}

// CreateToken creates a new paseto token.
func (p *Paseto) CreateToken(payload *user.TokenPayload) (string, error) {
	err := p.token.Set("payload", payload)
	if err != nil {
		return "", user.ErrTokenCreationFail
	}

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(p.duration)

	p.token.SetIssuedAt(issuedAt)
	p.token.SetNotBefore(issuedAt)
	p.token.SetExpiration(expiredAt)

	token := p.token.V4Encrypt(*p.key, nil)

	return token, nil
}

// VerifyToken verifies the paseto token.
func (pt *Paseto) VerifyToken(token string) (*user.TokenPayload, error) {
	var payload *user.TokenPayload

	parsedToken, err := pt.parser.ParseV4Local(*pt.key, token, nil)
	if err != nil {
		if err.Error() == "this token has expired" {
			return nil, user.ErrTokenExpired
		}
		return nil, user.ErrTokenInvalid
	}

	err = parsedToken.Get("payload", &payload)
	if err != nil {
		return nil, user.ErrTokenInvalid
	}

	return payload, nil
}
