package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
