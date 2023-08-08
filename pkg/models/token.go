package models

// Token represents a JWT token.
// It contains the following fields:
// - Token: the JWT token
type Token struct {
	Token string `json:"token"`
}
