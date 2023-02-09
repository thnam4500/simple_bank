package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// Create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check if the input token is valid
	VerifyToken(token string) (*Payload, error)
}
