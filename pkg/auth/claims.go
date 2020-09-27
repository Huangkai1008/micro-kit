package auth

// Claims is the interface for all the claims that can be made by the auth package.
type Claims interface {
	// HasExpired returns true if the claims have expired.
	HasExpired() bool
}
