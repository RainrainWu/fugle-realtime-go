package config

import (
	"os"
)

type FugleConfig interface {
	GetAPIToken() string
}

type fugleConfig struct {
	apiToken string
}

type FugleConfigOption interface {
	apply(*fugleConfig)
}

func NewFugleConfig(opts ...FugleConfigOption) FugleConfig {
	instance := &fugleConfig{
		apiToken: os.Getenv("FUGLE_API_TOKEN"),
	}
	for _, opt := range opts {
		opt.apply(instance)
	}
	return instance
}

func (conf *fugleConfig) GetAPIToken() string {
	return conf.apiToken
}
