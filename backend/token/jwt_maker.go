package token

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// ensure that the secret key is at least 32 characters long
const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, jwt.ErrInvalidKey
	}
	return &JWTMaker{secretKey}, nil
}

// create a new token for a specific username and duration
func (maker JWTMaker) CreateToken(userID pgtype.UUID, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// Verify Token checks if the token is valid or not
func (maker JWTMaker) VerifyToken(token string) (*Payload, error) {
	// define a key validation function
	keyFunc := func(token *jwt.Token) (any, error) {
		// type assertion: since we are using the HMAC algo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(maker.secretKey), nil
	}

	// parse the token
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, jwt.ErrTokenExpired
		} else if strings.Contains(err.Error(), "signature is invalid") {
			return nil, jwt.ErrSignatureInvalid
		}
		return nil, jwt.ErrInvalidKey
	}

	// get payload
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, jwt.ErrInvalidType
	}
	return payload, nil
}
