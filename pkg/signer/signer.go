package signer

import (
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword encrypts the given password using bcrypt.
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
