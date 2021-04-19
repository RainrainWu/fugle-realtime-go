package config

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	Config ConfigSet
)

type ConfigSet struct {
	Fugle FugleConfig
}

func NewConfigSet() ConfigSet {
	instance := ConfigSet{
		Fugle: NewFugleConfig(),
	}
	return instance
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Config = NewConfigSet()
}
