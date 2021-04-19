package main

import (
	"fmt"
	"log"

	"github.com/RainrainWu/fugle-realtime-go/client"
	"github.com/RainrainWu/fugle-realtime-go/config"
)

func main() {
	fmt.Println(config.Config.GetFugleConfig().GetAPIToken())

	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	if err != nil {
		log.Fatal("failed to init fugle api client")
	}
	result := myClient.Chart("2317", false)
	fmt.Println(result.APIVersion)
	fmt.Println(result.Data.Info.LastUpdatedAt)
}
