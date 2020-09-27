package consul

type Option func(*Options)

type Options struct {
	// Addr is the address of the consul server.
	Addr string

	// EnableHealthCheck enables the health check.
	EnableHealthCheck bool

	// HealthCheckInterval is the interval of the health check.
	HealthCheckInterval int

	// DeregisterCriticalServiceAfter is the time to wait before deregister a critical service.
	DeregisterCriticalServiceAfter int

	// HeartBeat enables the heart beat.
	HeartBeat bool
}

// WithAddr sets the address of the consul server.
func WithAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

// WithEnableHealthCheck enables the health check.
func WithEnableHealthCheck(enable bool) Option {
	return func(o *Options) {
		o.EnableHealthCheck = enable
	}
}

// WithHealthCheckInterval sets the interval of the health check.
func WithHealthCheckInterval(interval int) Option {
	return func(o *Options) {
		o.HealthCheckInterval = interval
	}
}

// WithDeregisterCriticalServiceAfter sets the time to wait before deregister a critical service.
func WithDeregisterCriticalServiceAfter(deregisterCriticalServiceAfter int) Option {
	return func(o *Options) {
		o.DeregisterCriticalServiceAfter = deregisterCriticalServiceAfter
	}
}

// WithHeartBeat enables the heart beat.
func WithHeartBeat(enable bool) Option {
	return func(o *Options) {
		o.HeartBeat = enable
	}
}
