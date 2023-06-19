package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(ss string, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(exp).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(ss))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
