package token

import "time"

type Maker interface {
	// create a new token for a specific username and duration
	CreateToken(userEmail string, duration time.Duration) (string, error)

	// Verify Token checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
