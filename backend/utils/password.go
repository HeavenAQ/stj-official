package utils

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// for decoding the password coming from the client
func B64Decode(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// for hashing the password before saving it to the database
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash password: %v", err)
	}
	return string(hashedPassword), nil
}

// check if the password is correct
func CheckPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("password is incorrect: %v", err)
	}
	return nil
}
