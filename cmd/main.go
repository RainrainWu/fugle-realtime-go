package main

import (
	"log"

	"github.com/RainrainWu/fugle-realtime-go/client"
	"github.com/RainrainWu/fugle-realtime-go/config"
)

func main() {
	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	if err != nil {
		log.Fatal("failed to init fugle api client")
	}
	result := myClient.Dealts("2330", false)
	result.PrettyPrint()
}
