package validator

type Option func(*Options)

// Options for the validator.
type Options struct {
	// The locale to use for validation messages.
	Locale string
}
