package registry

import "context"

type Registrar interface {
	// Register a service with the registry.
	Register(ctx context.Context, service *ServiceInstance) error
	// Deregister a service with the registry.
	Deregister(ctx context.Context, service *ServiceInstance) error
}

// ServiceInstance is an instance of a service in a discovery system.
type ServiceInstance struct {
	// ID is the unique instance ID as registered.
	ID string `json:"id"`

	// Name is the service name as registered.
	Name string `json:"name"`

	// Version is the version of the compiled.
	Version string `json:"version"`

	// Metadata is the kv pair metadata associated with the service instance.
	Metadata map[string]string `json:"metadata"`

	// Endpoints are the URLs that are available to the service instance.
	// schema:
	//   http://127.0.0.1:8000?isSecure=false
	//   grpc://127.0.0.1:9000?isSecure=false
	Endpoints []string `json:"endpoints"`
}
