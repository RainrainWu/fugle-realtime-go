package main

import (
	"fmt"

	"github.com/RainrainWu/fugle-realtime-go/config"
)

func main() {
	fmt.Println(config.Config.Fugle.GetAPIToken())
}
