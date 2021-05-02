package main

import (
	"github.com/RainrainWu/fugle-realtime-go/client"
	"github.com/RainrainWu/fugle-realtime-go/config"
	"github.com/sirupsen/logrus"
)

func main() {
	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	result, err := myClient.Meta("0056", false)
	if err != nil {
		logrus.Warn(err.Error())
	}
	result.PrettyPrint()
}
