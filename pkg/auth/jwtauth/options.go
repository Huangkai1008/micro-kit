package jwtauth

import (
	"time"
)

type Option func(*Options)

// Options is a struct for specifying configuration options for the JWT.
type Options struct {
	// SecretKey is the secret key used to sign the token.
	// It is required.
	SecretKey string

	// Issuer is the issuer of the token.
	Issuer string

	// Subject is the subject of the token.
	Subject string

	// Audience is the audience of the token.
	Audience []string

	// AccessTokenExpires is the expiration time of the access token.
	AccessTokenExpires time.Duration

	// RefreshTokenExpires is the expiration time of the refresh token.
	RefreshTokenExpires time.Duration
}

// WithIssuer sets the issuer of the token.
func WithIssuer(issuer string) Option {
	return func(o *Options) {
		o.Issuer = issuer
	}
}

// WithSubject sets the subject of the token.
func WithSubject(subject string) Option {
	return func(o *Options) {
		o.Subject = subject
	}
}

// WithAudience sets the audience of the token.
func WithAudience(audience ...string) Option {
	return func(o *Options) {
		o.Audience = audience
	}
}

// WithAccessTokenExpires sets the expiration time of the access token.
func WithAccessTokenExpires(expires time.Duration) Option {
	return func(o *Options) {
		o.AccessTokenExpires = expires
	}
}

// WithRefreshTokenExpires sets the expiration time of the refresh token.
func WithRefreshTokenExpires(expires time.Duration) Option {
	return func(o *Options) {
		o.RefreshTokenExpires = expires
	}
}
