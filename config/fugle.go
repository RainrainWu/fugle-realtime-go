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

type FugleOption interface {
	apply(*fugleConfig)
}

type fugleOptionFunc func(*fugleConfig)

func (f fugleOptionFunc) apply(c *fugleConfig) {
	f(c)
}

func NewFugleConfig(opts ...FugleOption) FugleConfig {
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
