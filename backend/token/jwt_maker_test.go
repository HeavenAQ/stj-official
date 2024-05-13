package token

import (
	"stj-ecommerce/utils"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	// create a new JWT maker
	maker, err := NewJWTMaker(utils.RandomAlphabetString(32))
	require.NoError(t, err)

	// payload information
	userID := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// Create a token
	token, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// verify the token
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	// create a new JWT maker
	maker, err := NewJWTMaker(utils.RandomAlphabetString(32))
	require.NoError(t, err)

	// create a token with a duration of -1
	userID := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	token, err := maker.CreateToken(userID, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// verify the token
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
	require.EqualError(t, err, jwt.ErrTokenExpired.Error())
}

func TestInvalidTokenAlgNone(t *testing.T) {
	// create a new JWT maker
	userID := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
	payload, err := NewPayload(userID, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	// create a token with the none algorithm
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	// create a new JWT maker
	maker, err := NewJWTMaker(utils.RandomAlphabetString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, jwt.ErrSignatureInvalid.Error())
	require.Nil(t, payload)

}
