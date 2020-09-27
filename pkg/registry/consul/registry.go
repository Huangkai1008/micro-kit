package consul

import (
	"context"
	"sync"

	"github.com/Huangkai1008/micro-kit/pkg/registry"
)

// Registry is a consul registry.
type Registry struct {
	client *Client
	lock   sync.RWMutex
}

// New returns a new consul registry.
func New(client *Client) *Registry {
	return &Registry{
		client: client,
	}
}

// Register a service instance.
func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	return r.client.Register(ctx, service)
}

// Deregister a service instance.
func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	return r.client.Deregister(ctx, service)
}
