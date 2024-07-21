package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Param func(*ConfigOptions)

type ConfigOptions struct {
	paths       []string
	name        string
	typeConfig  string
	readEnv     bool
	envOverride bool
}

func defaultOptions() *ConfigOptions {
	return &ConfigOptions{
		paths:       []string{"."},
		name:        "config",
		typeConfig:  "yml",
		readEnv:     true,
		envOverride: true,
	}
}

// New creates a new ConfigOptions.
func New(opts ...Param) *ConfigOptions {
	cfg := defaultOptions()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func (c *ConfigOptions) Load() (*ConfigOptions, error) {
	viper.SetConfigName(c.name)
	for _, path := range c.paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigType(c.typeConfig)
	if c.readEnv {
		viper.AutomaticEnv()
	}
	if c.envOverride {
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ConfigOptions) Unmarshal(rawVal any) error {
	return viper.Unmarshal(rawVal)
}
