package error

var (
	ErrAccountEmptyAuthHeader   = NewBadRequestError("Missing token in request header")
	ErrAccountInvalidAuthHeader = NewBadRequestError("Malformed token in request header")

	ErrInvalidToken = NewUnauthorizedError("The token is invalid")
	ErrTokenExpired = NewUnauthorizedError("The token has expired")
)
