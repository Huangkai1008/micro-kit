package consul

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	kitmsg "github.com/Huangkai1008/kit/pkg/message"
)

type Options struct {
	Addr                           string `mapstructure:"addr"`
	EnableHealthCheck              bool   `mapstructure:"enable_health_check"`
	HealthCheckInterval            int    `mapstructure:"health_check_interval"`
	DeregisterCriticalServiceAfter int    `mapstructure:"deregister_critical_service_after"`
	HeartBeat                      bool   `mapstructure:"heart_beat"`
}

// NewOptions creates a new set of o for the consul client.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.Sub("consul").Unmarshal(o); err != nil {
		return nil, errors.Wrap(err, kitmsg.LoadConfigError)
	}
	return o, err
}
