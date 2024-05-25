package utils

import (
	"time"
	"os"
	"github.com/dgrijalva/jwt-go"
	"log"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error generating JWT:", err)
		return "", err
	}
	log.Println("JWT generated successfully")
	return signedToken, nil
}

func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		log.Println("Error validating JWT:", err)
		return "", err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		log.Println("Invalid JWT token")
		return "", err
	}
	log.Println("JWT token validated successfully")
	return claims.Email, nil
}
