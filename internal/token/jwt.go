package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/golang-jwt/jwt/v5"
)

// jwtTokenMaker implements [user.Tokenizer] interface.
type jwtTokenMaker struct {
	key string
}

// jwtClaims represents the JWT claims structure.
type jwtClaims struct {
	Payload user.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

// NewJWTTokenMaker creates a new jwtTokenMaker instance.
func NewJWTTokenMaker(key string) *jwtTokenMaker {
	return &jwtTokenMaker{
		key: key,
	}
}

// CreateToken generates a new JWT token with a payload and expiration duration.
func (j *jwtTokenMaker) CreateToken(
	payload *user.TokenPayload, duration time.Duration,
) (token string, expiredAt time.Time, err error) {
	issuedAt := time.Now()
	expiredAt = issuedAt.Add(duration)

	claims := jwtClaims{
		Payload: *payload,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			NotBefore: jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}

	tokenBuilder := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := tokenBuilder.SignedString([]byte(j.key))
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expiredAt, nil
}

// VerifyToken validates the JWT and extracts the payload.
func (j *jwtTokenMaker) VerifyToken(token string) (*user.TokenPayload, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&jwtClaims{},
		j.keyFunc,
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*jwtClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims.Payload, nil
}

// keyFunc returns the key for validating the JWT token.
func (jt *jwtTokenMaker) keyFunc(token *jwt.Token) (any, error) {
	if token.Method != jwt.SigningMethodHS256 {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte(jt.key), nil
}
