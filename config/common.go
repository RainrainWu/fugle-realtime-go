package config

import (
	"log"

	"github.com/joho/godotenv"
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
		log.Fatal("Error loading .env file")
		return
	}
	Config = NewConfigSet()
}
