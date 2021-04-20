package main

import (
	"github.com/RainrainWu/fugle-realtime-go/client"
	"github.com/RainrainWu/fugle-realtime-go/config"
	"github.com/RainrainWu/fugle-realtime-go/logger"
)

func main() {
	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	if err != nil {
		logger.PrintLogger.Fatal("failed to init fugle api client")
	}
	result := myClient.Dealts("2330", false)
	result.PrettyPrint()
}
