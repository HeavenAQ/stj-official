package utils

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func Test_B64Encode(t *testing.T) {
	input := "hello world"
	expectedOutput := base64.StdEncoding.EncodeToString([]byte(input))
	encodedString := B64Encode(input)
	require.Equal(t, expectedOutput, encodedString)
}

func Test_B64Decode(t *testing.T) {
	// valid base64 string
	validInput := B64Encode("hello world")
	expectedOutput := "hello world"
	decodedString, err := B64Decode(validInput)
	require.NoError(t, err)
	require.Equal(t, expectedOutput, decodedString)

	// invalid base64 string
	invalidInput := "not_base64^^"
	_, err = B64Decode(invalidInput)
	require.Error(t, err)
}

func Test_HashPassword(t *testing.T) {
	password := "testpassword123"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	// Ensure the hash is valid
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	require.NoError(t, err)
}

func Test_CheckPassword(t *testing.T) {
	password := "securePassword"
	hashedPassword, _ := HashPassword(password)

	// password is encoded with base64 to simulate the password coming from the client
	// Check with the correct password
	err := CheckPassword(hashedPassword, B64Encode(password))
	require.NoError(t, err)

	// Check with the incorrect password
	err = CheckPassword(hashedPassword, B64Encode("wrongPassword"))
	require.Error(t, err)
}
