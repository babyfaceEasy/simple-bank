package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


const (
	minSecretKeySize = 32
)

// JWTMaker is a JSON Web token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey: secretKey}, nil
}

// CreateToken create and sign a token for a specific username and duration.
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	claims := jwt.RegisteredClaims {
		ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
		ID: payload.ID.String(),
		IssuedAt: jwt.NewNumericDate(payload.IssuedAt),
		Issuer: payload.Username,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the input token is valid or not.
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		// verr, ok := err.(*jwt.Valida)
		return nil, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, ErrInvalidToken
	} 
	payload := &Payload{
		ID: id,
		Username: claims.Issuer,
		ExpiredAt: claims.ExpiresAt.Time,
		IssuedAt: claims.IssuedAt.Time,

	}

	return payload, nil
}