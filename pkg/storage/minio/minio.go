package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	DefaultEndpoint = "localhost:9000"
	DefaultRegion   = "cn-north-1"
)

// New returns new minioClient instance with options.
//
// The default options are:
// 	- Endpoint: DefaultEndpoint
// 	- UseSSL: false
// 	- Region: DefaultRegion
//
func New(opts ...Option) (*minio.Client, error) {
	o := Options{
		Endpoint: DefaultEndpoint,
		UseSSL:   false,
		Region:   DefaultRegion,
	}

	for _, opt := range opts {
		opt(&o)
	}

	minioClient, err := minio.New(o.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(o.AccessKeyID, o.SecretAccessKey, ""),
		Secure: o.UseSSL,
		Region: o.Region,
	})
	if err != nil {
		return nil, err
	}
	return minioClient, err
}
