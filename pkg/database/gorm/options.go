package gorm

type Option func(*Options)

type Options struct {
	// Source is the database connection string.
	Source string

	// MaxIdleConns is the maximum number of connections in the idle connection pool.
	MaxIdleConnections int

	// MaxOpenConns is the maximum number of open connections to the database.
	MaxOpenConnections int

	// EnableAutoMigrate is the flag to enable auto migrate.
	EnableAutoMigrate bool
}

// WithSource sets the database connection string.
func WithSource(source string) Option {
	return func(o *Options) {
		o.Source = source
	}
}

// WithMaxIdleConns sets the maximum number of connections in the idle connection pool.
func WithMaxIdleConns(maxIdleConns int) Option {
	return func(o *Options) {
		o.MaxIdleConnections = maxIdleConns
	}
}

// WithMaxOpenConns sets the maximum number of open connections to the database.
func WithMaxOpenConns(maxOpenConns int) Option {
	return func(o *Options) {
		o.MaxOpenConnections = maxOpenConns
	}
}

// WithEnableAutoMigrate sets the flag to enable auto migrate.
func WithEnableAutoMigrate(enableAutoMigrate bool) Option {
	return func(o *Options) {
		o.EnableAutoMigrate = enableAutoMigrate
	}
}
