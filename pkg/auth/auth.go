package auth

// Auth is the interface for authentication.
type Auth interface {
	// CreateAccessToken create a new token.
	// The identity of this token, which can be any data that is json serializable.
	CreateAccessToken(identity interface{}, fresh bool) (string, error)

	// CreateRefreshToken create a new refresh token.
	// The identity of this token, which can be any data that is json serializable.
	CreateRefreshToken(identity interface{}) (string, error)

	// ParseToken parse a token.
	ParseToken(token string) (Claims, error)
}
