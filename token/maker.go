package token

import (
	"time"
)

// Maker  is an interface for managing tokens.
type Maker interface {
	// CreateToken create and sign a token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the input token is valid or not.
	VerifyToken(token string) (*Payload, error)
}
