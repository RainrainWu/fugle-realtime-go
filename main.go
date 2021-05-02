package main

import (
	"fmt"

	"github.com/RainrainWu/fugle-realtime-go/client"
	"github.com/RainrainWu/fugle-realtime-go/config"
)

func main() {
	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	result, err := myClient.Meta("0056", false)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		result.PrettyPrint()
	}
}
