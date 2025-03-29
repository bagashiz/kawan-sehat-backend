package token

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

// pasetoTokenMaker implements [user.Tokenizer] interface
// and provides an access to the paseto library.
type pasetoTokenMaker struct {
	token  *paseto.Token
	key    *paseto.V4SymmetricKey
	parser *paseto.Parser
}

// NewPasetoTokenMaker creates a new pasetoTokenMaker instance.
func NewPasetoTokenMaker() *pasetoTokenMaker {
	token := paseto.NewToken()
	key := paseto.NewV4SymmetricKey()
	parser := paseto.NewParser()

	return &pasetoTokenMaker{
		&token,
		&key,
		&parser,
	}
}

// CreateToken creates a new paseto token.
func (p *pasetoTokenMaker) CreateToken(
	payload *user.TokenPayload, duration time.Duration,
) (token string, expiredAt time.Time, err error) {
	issuedAt := time.Now()
	expiredAt = issuedAt.Add(duration)

	p.token.SetIssuedAt(issuedAt)
	p.token.SetNotBefore(issuedAt)
	p.token.SetExpiration(expiredAt)

	if err := p.token.Set("payload", payload); err != nil {
		return "", time.Time{}, err
	}

	token = p.token.V4Encrypt(*p.key, nil)

	return token, expiredAt, nil
}

// VerifyToken verifies the paseto token.
func (pt *pasetoTokenMaker) VerifyToken(token string) (*user.TokenPayload, error) {
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
