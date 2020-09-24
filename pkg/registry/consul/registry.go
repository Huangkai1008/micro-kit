package consul

import (
	"context"
	"sync"

	"github.com/google/wire"

	"github.com/Huangkai1008/micro-kit/pkg/registry"
)

type Registry struct {
	client *Client
	lock   sync.RWMutex
}

func New(client *Client) *Registry {
	return &Registry{
		client: client,
	}
}

func (r *Registry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	return r.client.Register(ctx, service)
}

func (r *Registry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	return r.client.Deregister(ctx, service)
}

var ProviderSet = wire.NewSet(New, NewClient, NewOptions, wire.Bind(new(registry.Registrar), new(*Registry)))
