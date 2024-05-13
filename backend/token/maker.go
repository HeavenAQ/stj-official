package token

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Maker interface {
	// create a new token for a specific username and duration
	CreateToken(userID pgtype.UUID, duration time.Duration) (string, error)

	// Verify Token checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
