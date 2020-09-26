package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/Huangkai1008/micro-kit/pkg/message"
	"github.com/Huangkai1008/micro-kit/pkg/transport/http"
)

// Application is the application service.
// It contains all the services and configurations.
type Application struct {
	Name       string
	Version    string
	logger     *zap.Logger
	httpServer *http.Server
	grpcServer *grpc.Server
	minioCli   *minio.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

// Option is the option of application.
type Option func(app *Application) error

// New returns a new Application.
func New(name, version string, logger *zap.Logger, options ...Option) (*Application, error) {

	a := &Application{
		Name:    name,
		Version: version,
		logger:  logger.With(zap.String("type", "Application")),
	}

	for _, option := range options {
		if err := option(a); err != nil {
			return nil, err
		}
	}

	return a, nil
}

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

// Start Application.
func (a *Application) Start() error {
	if err := a.httpServer.Start(); err != nil {
		return errors.WithMessage(err, message.HTTPServerStartError)
	}
	a.logger.Info("Application Started", zap.String("name", a.Name))
	return nil
}

// Stop Application.
func (a *Application) Stop() error {
	if err := a.httpServer.Stop(); err != nil {
		return errors.WithMessage(err, message.HTTPServerStopError)
	}
	a.logger.Info("Server exiting ...")
	return nil
}

// AwaitSignal await the signal to stop the server
func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	a.logger.Info("Receive a signal", zap.String("signal", s.String()))
	if a.httpServer != nil {
		if err := a.Stop(); err != nil {
			a.logger.Warn("Stop HTTP server error", zap.Error(err))
		}
	}
	os.Exit(0)
}
