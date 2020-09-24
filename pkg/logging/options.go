package logging

import "github.com/spf13/viper"

type Options struct {
	Level    int
	FileName string
	Stdout   bool
}

// NewOptions returns new log options.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}
	return o, err
}
