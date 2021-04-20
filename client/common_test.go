package client_test

import (
	"testing"

	"github.com/RainrainWu/fugle-realtime-go/client"
	"github.com/RainrainWu/fugle-realtime-go/config"
	"github.com/stretchr/testify/assert"
)

func TestClientChartAPI(t *testing.T) {

	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	assert.Nil(t, err)

	chart := myClient.Chart("2330", false)
	assert.NotNil(t, chart.Data.Chart)
}

func TestClientQuoteAPI(t *testing.T) {

	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	assert.Nil(t, err)

	result := myClient.Quote("2330", false)
	assert.NotNil(t, result.Data.Quote.Order.Bestasks)
	assert.NotNil(t, result.Data.Quote.Order.Bestbids)
}

func TestClientMetaAPI(t *testing.T) {

	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	assert.Nil(t, err)

	result := myClient.Meta("2330", false)
	assert.NotEmpty(t, result.Data.Meta.Namezhtw)
	assert.NotEmpty(t, result.Data.Meta.Industryzhtw)
	assert.NotEmpty(t, result.Data.Meta.Volumeperunit)
	assert.NotEmpty(t, result.Data.Meta.Currency)
	assert.NotEmpty(t, result.Data.Meta.Typezhtw)
	assert.NotEmpty(t, result.Data.Meta.Abnormal)
}

func TestClientDealtsAPI(t *testing.T) {

	myClient, err := client.NewFugleClient(
		client.ConfigOption(config.Config),
	)
	assert.Nil(t, err)

	result := myClient.Dealts("2330", false)
	assert.NotEmpty(t, result.Data.Dealts)
}
