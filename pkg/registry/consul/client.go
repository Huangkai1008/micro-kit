package consul

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"

	"github.com/Huangkai1008/micro-kit/pkg/registry"
)

type Client struct {
	cli    *api.Client
	ctx    context.Context
	cancel context.CancelFunc
	logger *zap.Logger

	*Options
}

func NewClient(o *Options, logger *zap.Logger) (*Client, error) {
	cli, err := api.NewClient(&api.Config{
		Address: o.Addr,
	})
	if err != nil {
		return nil, err
	}

	c := &Client{
		cli:     cli,
		logger:  logger,
		Options: o,
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c, nil
}

func (c *Client) Register(_ context.Context, service *registry.ServiceInstance) (err error) {
	addresses := make(map[string]api.ServiceAddress)
	checkAddresses := make([]string, 0, len(service.Endpoints))
	for _, endpoint := range service.Endpoints {
		raw, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		addr := raw.Hostname()
		port, _ := strconv.ParseUint(raw.Port(), 10, 16)

		checkAddresses = append(checkAddresses, net.JoinHostPort(addr, strconv.FormatUint(port, 10)))
		addresses[raw.Scheme] = api.ServiceAddress{Address: endpoint, Port: int(port)}
	}
	asr := &api.AgentServiceRegistration{
		ID:              service.ID,
		Name:            service.Name,
		Meta:            service.Metadata,
		Tags:            []string{fmt.Sprintf("version=%s", service.Version)},
		TaggedAddresses: addresses,
	}
	if len(checkAddresses) > 0 {
		host, portRaw, _ := net.SplitHostPort(checkAddresses[0])
		port, _ := strconv.ParseInt(portRaw, 10, 32)
		asr.Address = host
		asr.Port = int(port)
	}

	if c.EnableHealthCheck {
		for _, address := range checkAddresses {
			asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
				TCP:                            address,
				Interval:                       fmt.Sprintf("%ds", c.HealthCheckInterval),
				DeregisterCriticalServiceAfter: fmt.Sprintf("%dm", c.DeregisterCriticalServiceAfter),
				Timeout:                        "5s",
			})
		}
	}

	if c.HeartBeat {
		asr.Checks = append(asr.Checks, &api.AgentServiceCheck{
			CheckID:                        "service:" + service.ID,
			TTL:                            fmt.Sprintf("%ds", c.HealthCheckInterval*2),
			DeregisterCriticalServiceAfter: fmt.Sprintf("%dm", c.DeregisterCriticalServiceAfter),
		})
	}

	err = c.cli.Agent().ServiceRegister(asr)
	if err != nil {
		return err
	}
	if c.HeartBeat {
		go func() {
			time.Sleep(time.Second)
			err = c.cli.Agent().UpdateTTL("service:"+service.ID, "pass", "pass")
			if err != nil {
				c.logger.Error("Update ttl heartbeat to consul failed!", zap.Error(err))
			}
			ticker := time.NewTicker(time.Second * time.Duration(c.HealthCheckInterval))
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					err = c.cli.Agent().UpdateTTL("service:"+service.ID, "pass", "pass")
					if err != nil {
						c.logger.Error("Update ttl heartbeat to consul failed!", zap.Error(err))
					}
				case <-c.ctx.Done():
					return
				}
			}
		}()
	}
	return nil
}

func (c *Client) Deregister(_ context.Context, service *registry.ServiceInstance) error {
	c.cancel()
	return c.cli.Agent().ServiceDeregister(service.ID)
}
