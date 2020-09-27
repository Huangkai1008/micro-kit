// Package http implements the HTTP transport protocol.
package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"go.uber.org/zap"

	"github.com/Huangkai1008/micro-kit/pkg/registry"
	"github.com/Huangkai1008/micro-kit/pkg/util"
)

var (
	DefaultReadTimeout  = 60 * time.Second
	DefaultWriteTimeout = 60 * time.Second
	DefaultMode         = DebugMode
)

// Server is the HTTP server.
type Server struct {
	Name       string
	Version    string
	httpServer *http.Server
	router     http.Handler
	logger     *zap.Logger
	registrar  registry.Registrar

	*Options
}

// New creates a new HTTP server.
//
// The default options are:
//  - ReadTimeout: DefaultReadTimeout
//  - WriteTimeout: DefaultWriteTimeout
//  - Mode: DebugMode
//
func New(logger *zap.Logger, router http.Handler, registrar registry.Registrar, opts ...Option) *Server {
	o := Options{
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		Mode:         DefaultMode,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &Server{
		httpServer: &http.Server{
			Addr:           o.Addr(),
			ReadTimeout:    o.ReadTimeout * time.Second,
			WriteTimeout:   o.WriteTimeout * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		router:    router,
		logger:    logger.With(zap.String("type", "http.Server")),
		registrar: registrar,
		Options:   &o,
	}
}

// Start http server.
func (s *Server) Start() error {
	s.httpServer.Handler = s.router
	s.logger.Info("HTTP server starting ...", zap.String("addr", s.Addr()))

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Start HTTP server error", zap.Error(err))
		}
	}()

	if err := s.register(s.registrar); err != nil {
		return errors.WithMessage(err, "Register HTTP service error")
	}
	return nil
}

// Stop http server.
func (s *Server) Stop() error {
	s.logger.Info("HTTP server stopping ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.deregister(s.registrar); err != nil {
		return errors.WithMessage(err, "Deregister HTTP service error")
	}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.WithMessage(err, "Shutdown HTTP server error")
	}
	return nil
}

// GetID returns the unique identifier of the server.
// It is useful when a service starts more than one server.
func (s *Server) GetID() string {
	return fmt.Sprintf("%s[%s]", s.Name, s.IntranetAddr())
}

func (s *Server) Service() *registry.ServiceInstance {
	id := s.GetID()
	return &registry.ServiceInstance{
		ID:       id,
		Name:     s.Name,
		Version:  s.Version,
		Metadata: map[string]string{"service": s.Name, "version": s.Version},
		Endpoints: []string{
			fmt.Sprintf("http://%s?isSecure=false", s.IntranetAddr()),
		},
	}
}

func (s *Server) IntranetAddr() string {
	return fmt.Sprintf("%s:%d", util.GetIntranetIP(), s.Port)
}

func (s *Server) register(registry registry.Registrar) error {
	return registry.Register(context.Background(), s.Service())
}

func (s *Server) deregister(registry registry.Registrar) error {
	return registry.Deregister(context.Background(), s.Service())
}
