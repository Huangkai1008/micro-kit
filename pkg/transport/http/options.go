package http

import (
	"fmt"
	"time"
)

const (
	DebugMode   = "debug"
	TestMode    = "testing"
	ReleaseMode = "release"
)

type Option func(*Options)

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

// Addr returns the address of the HTTP server.
// The format is "host:port".
func (o *Options) Addr() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

// WithHost sets the hostname of the HTTP server.
func WithHost(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

// WithPort sets the port of the HTTP server.
func WithPort(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

// WithReadTimeout sets the read-timeout for the HTTP server.
func WithReadTimeout(readTimeout time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = readTimeout
	}
}

// WithWriteTimeout sets the write-timeout for the HTTP server.
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = writeTimeout
	}
}

// WithMode sets the mode for the HTTP server.
func WithMode(mode string) Option {
	return func(o *Options) {
		o.Mode = mode
	}
}
