package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto *paseto.V2
	symmeticKey []byte
}

func NewPasetoMaker(symmeticKey string) (Maker, error) {
	if len(symmeticKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size, it must have exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symmeticKey: []byte(symmeticKey),
	}

	return maker, nil
}

// // CreateToken create and sign a token for a specific username and duration.
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	
	return maker.paseto.Encrypt(maker.symmeticKey, payload, nil)
}

// VerifyToken checks if the input token is valid or not.
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmeticKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}

