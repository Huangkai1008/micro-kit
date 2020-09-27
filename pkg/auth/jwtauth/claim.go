package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenType int

const (
	AccessTokenType TokenType = iota + 1
	RefreshTokenType
)

// Claims is the custom claims type.
type Claims struct {
	// The identity of this token, which can be any data that is json serializable.
	Identity interface{}

	// Every time a user authenticates by providing a username and password,
	// they receive a Fresh access token that can access any route.
	//
	// But after some time, that token should no longer be considered fresh,
	// and some critical or dangerous routes will be blocked until the user verifies their password again.
	Fresh bool

	// TokenType includes AccessTokenType and RefreshTokenType.
	TokenType TokenType

	jwt.RegisteredClaims
}

// HasExpired returns true if the token has expired.
func (c *Claims) HasExpired() bool {
	return c.VerifyExpiresAt(time.Now(), true)
}
