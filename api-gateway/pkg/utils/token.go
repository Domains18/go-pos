package utils

import (
"time"

"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mike")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Creates and signs a new JWT token for a given username

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Notiwa",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

