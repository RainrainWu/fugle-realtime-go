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
		logrus.Fatal("failed to init fugle api client")
	}
	result := myClient.Meta("0056", false)
	result.PrettyPrint()
}
