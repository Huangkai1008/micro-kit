package auth

// Claims is the interface for all the claims that can be made by the auth package.
type Claims interface {
	// GetIdentity return the identity of this token, which can be any data that is json serializable.
	GetIdentity() interface{}

	// HasExpired returns true if the claims have expired.
	HasExpired() bool
}
