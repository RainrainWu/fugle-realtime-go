# fugle-realtime-go

[![Build Status](https://circleci.com/gh/RainrainWu/fugle-realtime-go.svg?style=shield)](https://app.circleci.com/pipelines/github/RainrainWu/fugle-realtime-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/RainrainWu/fugle-realtime-go)](https://goreportcard.com/report/github.com/RainrainWu/fugle-realtime-go)
[![codecov](https://codecov.io/gh/RainrainWu/fugle-realtime-go/branch/main/graph/badge.svg?token=4E4PS8S4TX)](https://codecov.io/gh/RainrainWu/fugle-realtime-go)
[![Code Climate](https://codeclimate.com/github/RainrainWu/fugle-realtime-go/badges/gpa.svg)](https://codeclimate.com/github/RainrainWu/fugle-realtime-go)

Fugle Realtime Go is a go package to query realtime stock data of Taiwan market through API provided by [Fugle API](https://developer.fugle.tw/)

# Getting Started

## Install the package
```bash
$ go get github.com/RainrainWu/fugle-realtime-go
```

## Set up your Fugle API token
```bash
$ export FUGLE_API_TOKEN=<YOUR_TOKEN>
```

## Quick demo
```go
package main

import (
	"log"

	"github.com/RainrainWu/fugle-realtime-go/client"
)

func main() {

	myClient, err := client.NewFugleClient()
	if err != nil {
		log.Fatal("failed to init fugle api client")
	}
	result := myClient.Meta("2330", false)
	result.PrettyPrint()
}

```

## Functions
### func (client.FugleClient).Chart(symbolID string, oddLot bool) client.FugleAPIResponse

Access the [Chart API](https://developer.fugle.tw/document/intraday/chart)

### func (client.FugleClient).Quote(symbolID string, oddLot bool) client.FugleAPIResponse

Access the [Quote API](https://developer.fugle.tw/document/intraday/quote)

### func (client.FugleClient).Meta(symbolID string, oddLot bool) client.FugleAPIResponse

Access the [Meta API](https://developer.fugle.tw/document/intraday/meta)

### func (client.FugleClient).Dealts(symbolID string, oddLot bool) client.FugleAPIResponse

Access the [Dealts API](https://developer.fugle.tw/document/intraday/dealts)
