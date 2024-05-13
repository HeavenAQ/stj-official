package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Payload struct {
	UserID pgtype.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func NewPayload(UserID pgtype.UUID, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		UserID,
		jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "stj-ecommerce",
			Subject:   "auth",
			Audience:  []string{"stj-ecommerce"},
		},
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt.Time) {
		return jwt.ErrTokenExpired
	}
	return nil
}
