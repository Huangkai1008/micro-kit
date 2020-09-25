package gorm

import (
	"fmt"

	"github.com/spf13/viper"
)

type Options struct {
	User       string
	Password   string
	Host       string
	Port       int
	DBName     string
	Parameters string

	MaxIdleConnections int
	MaxOpenConnections int
	EnableAutoMigrate  bool
}

// NewOptions returns new log options.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("database", o); err != nil {
		return nil, err
	}

	return o, err
}

// DSN returns data source name.
func (o *Options) DSN() string {
	const dsn = "%s:%s@tcp(%s:%d)/%s?%s"
	return fmt.Sprintf(dsn, o.User, o.Password, o.Host, o.Port, o.DBName, o.Parameters)
}
