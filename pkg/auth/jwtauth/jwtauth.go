package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/Huangkai1008/micro-kit/pkg/auth"
)

var (
	DefaultAccessTokenExpires  = time.Hour * 2
	DefaultRefreshTokenExpires = time.Hour * 24 * 30
	DefaultSigningMethod       = jwt.SigningMethodHS256
)

type JwtAuth struct {
	*Options
}

// New returns new JwtAuth.
//
// The SecretKey is used to sign the token.
//
// The default options are:
//  - AccessTokenExpires: DefaultAccessTokenExpires
//  - RefreshTokenExpires: DefaultRefreshTokenExpires
//  - SigningMethod: DefaultSigningMethod
//
func New(secretKey string, opts ...Option) *JwtAuth {
	o := Options{
		SecretKey:           secretKey,
		AccessTokenExpires:  DefaultAccessTokenExpires,
		RefreshTokenExpires: DefaultRefreshTokenExpires,
	}

	for _, opt := range opts {
		opt(&o)
	}
	return &JwtAuth{&o}
}

// CreateAccessToken create a new access token.
// The identity of this token, which can be any data that is json serializable.
func (j *JwtAuth) CreateAccessToken(identity interface{}, fresh bool) (string, error) {
	claims := Claims{
		Identity:  identity,
		Fresh:     fresh,
		TokenType: AccessTokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			Subject:   j.Subject,
			Audience:  j.Audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessTokenExpires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(DefaultSigningMethod, claims)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	return tokenString, err
}

// CreateRefreshToken create a new refresh token.
// The identity of this token, which can be any data that is json serializable.
func (j *JwtAuth) CreateRefreshToken(identity interface{}) (string, error) {
	claims := Claims{
		Identity:  identity,
		Fresh:     false,
		TokenType: RefreshTokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			Subject:   j.Subject,
			Audience:  j.Audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshTokenExpires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(DefaultSigningMethod, claims)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	return tokenString, err
}

// ParseToken parse token string to Claims.
func (j *JwtAuth) ParseToken(tokenString string) (auth.Claims, error) {
	return j.ParseJwtToken(tokenString)
}

func (j *JwtAuth) ParseJwtToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
