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

// CheckPasswordHash compares a hashed password with plaintext password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Password comparison failed:", err)
		return false
	}
	return true
}