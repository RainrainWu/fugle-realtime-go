package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	Config ConfigSet
)

type ConfigSet interface {
	GetFugleConfig() FugleConfig
}

type configSet struct {
	fugle FugleConfig
}

func NewConfigSet() ConfigSet {
	instance := &configSet{
		fugle: NewFugleConfig(),
	}
	return instance
}

func (conf *configSet) GetFugleConfig() FugleConfig {
	return conf.fugle
}

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn(
			"error loading .env file, current environment " +
				"variables would be used directly",
		)
	}
	Config = NewConfigSet()
}
