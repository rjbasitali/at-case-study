package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// NewToken generates a new JWT token.
// It accepts a secret string and an expiry duration.
// It returns the token as a string and an error if the token could not be generated.
// The token is signed using the secret and expires after the specified duration.
// The token contains the following claims:
// - nbf: the time when the token was issued
// - iat: the time when the token was issued
// - exp: the time when the token will expire
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
