package server

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(username string, expireDuration time.Duration) (string, error) {
	mySigningKey := []byte("ct-secret-key")

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		Issuer:    "ct-backend-course",
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	mySigningKey := []byte("ct-secret-key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
