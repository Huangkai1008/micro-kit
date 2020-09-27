package minio

type Option func(*Options)

// Options for MinIO storage.
type Options struct {
	// Endpoint is the endpoint of the MinIO storage.
	Endpoint string

	// AccessKeyID is the access key ID of the MinIO storage.
	AccessKeyID string

	// SecretAccessKey is the secret access key of the MinIO storage.
	SecretAccessKey string

	// UseSSL specifies whether to use SSL when accessing the MinIO storage.
	UseSSL bool

	// UseProxy specifies whether to use proxy when accessing the MinIO storage.
	Region string
}

// WithEndpoint sets the endpoint of the MinIO storage.
func WithEndpoint(endpoint string) Option {
	return func(o *Options) {
		o.Endpoint = endpoint
	}
}

// WithAccessKeyID sets the access key ID of the MinIO storage.
func WithAccessKeyID(accessKeyID string) Option {
	return func(o *Options) {
		o.AccessKeyID = accessKeyID
	}
}

// WithSecretAccessKey sets the secret access key of the MinIO storage.
func WithSecretAccessKey(secretAccessKey string) Option {
	return func(o *Options) {
		o.SecretAccessKey = secretAccessKey
	}
}

// WithUseSSL specifies whether to use SSL when accessing the MinIO storage.
func WithUseSSL(useSSL bool) Option {
	return func(o *Options) {
		o.UseSSL = useSSL
	}
}

// WithRegion specifies whether to use proxy when accessing the MinIO storage.
func WithRegion(region string) Option {
	return func(o *Options) {
		o.Region = region
	}
}
