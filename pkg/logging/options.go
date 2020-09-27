package logging

type Option func(*Options)

type Options struct {
	// FileName is log file name.
	// It is required.
	FileName string

	// Level is log level.
	Level int

	// Stdout is whether enable stdout.
	Stdout bool
}

// WithLevel sets the log level.
func WithLevel(level int) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// WithStdout sets the flag to enable stdout.
func WithStdout(stdout bool) Option {
	return func(o *Options) {
		o.Stdout = stdout
	}
}
