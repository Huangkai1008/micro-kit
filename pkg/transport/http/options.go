package http

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	kitmsg "github.com/Huangkai1008/kit/pkg/message"
)

const (
	DebugMode   = "debug"
	TestMode    = "testing"
	ReleaseMode = "release"
)

// Options for the HTTP server.
type Options struct {
	// The hostname of the HTTP server.
	Host string

	// The port of the HTTP server.
	Port int

	// Read timeout for the HTTP server.
	ReadTimeout time.Duration

	// Write timeout for the HTTP server.
	WriteTimeout time.Duration

	// Mode of the HTTP server. Includes DebugMode, TestMode, and ReleaseMode.
	Mode string
}

// NewOptions creates a new set of o for the HTTP server.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, errors.Wrap(err, kitmsg.LoadConfigError)
	}
	return o, err
}

// Addr returns the address of the HTTP server.
// The format is "host:port".
func (o *Options) Addr() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}
