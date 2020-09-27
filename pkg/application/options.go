package application

import (
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc"

	"github.com/Huangkai1008/micro-kit/pkg/transport/http"
)

// Option is the option of application.
type Option func(app *Application) error

// WithHttpServer sets the http server.
func WithHttpServer(s *http.Server) Option {
	return func(a *Application) error {
		s.Name = a.Name
		s.Version = a.Version
		a.httpServer = s
		return nil
	}
}

// WithGrpcServer sets the grpc server.
func WithGrpcServer(s *grpc.Server) Option {
	return func(a *Application) error {
		a.grpcServer = s
		return nil
	}
}

// WithMinioCli sets the minio client.
func WithMinioCli(c *minio.Client) Option {
	return func(a *Application) error {
		a.minioCli = c
		return nil
	}
}
