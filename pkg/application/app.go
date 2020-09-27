package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"google.golang.org/grpc"

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

// New returns a new Application.
func New(name, version string, logger *zap.Logger, opts ...Option) (*Application, error) {

	a := &Application{
		Name:    name,
		Version: version,
		logger:  logger.With(zap.String("type", "Application")),
	}

	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}

	return a, nil
}

// Start Application.
func (a *Application) Start() error {
	if err := a.httpServer.Start(); err != nil {
		return err
	}
	a.logger.Info("Application Started", zap.String("name", a.Name))
	return nil
}

// Stop Application.
func (a *Application) Stop() error {
	if err := a.httpServer.Stop(); err != nil {
		return err
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
