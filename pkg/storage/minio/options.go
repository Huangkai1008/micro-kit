package minio

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Huangkai1008/kit/pkg/message"
)

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

// NewOptions creates a new set of o for the minio client.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("minio", o); err != nil {
		return nil, errors.Wrap(err, message.LoadConfigError)
	}
	return o, err
}
