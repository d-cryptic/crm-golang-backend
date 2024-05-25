package utils

import (
	"log"
	"encoding/base64"
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password with a custom salt
func HashPassword(password string) (string, error) {
    // Generate a random salt
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }

    // Concatenate the password and salt
    saltedPassword := append([]byte(password), salt...)

    // Generate the hash with the salted password
    hashedBytes, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    // Encode the hash to a base64 string
    hashedPassword := base64.StdEncoding.EncodeToString(hashedBytes)
    return hashedPassword, nil
}

// CheckPasswordHash compares a hashed password with plaintext password
func CheckPasswordHash(password, hash string) bool {
    // Decode the hash from base64
    hashedBytes, err := base64.StdEncoding.DecodeString(hash)
    if err != nil {
        log.Println("Error decoding hash:", err)
        return false
    }

    // Compare the hashed password with the provided password
    err = bcrypt.CompareHashAndPassword(hashedBytes, []byte(password))
    if err != nil {
        log.Println("Password comparison failed:", err)
        return false
    }
    return true
}
