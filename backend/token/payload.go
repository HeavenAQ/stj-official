package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

func NewPayload(userEmail string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		userEmail,
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
