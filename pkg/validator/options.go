package validator

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	kitmsg "github.com/Huangkai1008/kit/pkg/message"
)

// Options for the validator.
type Options struct {
	// The locale to use for validation messages.
	Locale string
}

// NewOptions creates a new set of o for the HTTP server.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, kitmsg.LoadConfigError)
	}
	return o, err
}
