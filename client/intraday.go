package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RainrainWu/fugle-realtime-go/config"
)

type FugleAPIResponse struct {
	APIVersion string       `json:"apiVersion"`
	Data       FugleAPIData `json:"data"`
}

type FugleAPIData struct {
	Info  FugleAPIInfo             `json:"info"`
	Chart map[string]FugleAPIPrice `json:"chart"`
}

type FugleAPIInfo struct {
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	Date          string    `json:"date"`
	Mode          string    `json:"mode"`
	SymbolID      string    `json:"symbolId"`
	CountryCode   string    `json:"countryCode"`
	Timezone      string    `json:"timeZone"`
}

type FugleAPIPrice struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Unit   int     `json:"unit"`
	Volume int     `json:"volume"`
}

type FugleClient interface {
	Chart(symbolID string, oddLot bool) FugleAPIResponse
}

type fugleClient struct {
	host          string
	headerAccept  string
	chartEndpoint string
	config        config.ConfigSet
}

type FugleClientOption interface {
	apply(*fugleClient)
}

type fugleOptionFunc func(*fugleClient)

func (f fugleOptionFunc) apply(c *fugleClient) {
	f(c)
}

func ConfigOption(conf config.ConfigSet) FugleClientOption {
	return fugleOptionFunc(func(cli *fugleClient) {
		cli.config = conf
	})
}

func NewFugleClient(opts ...FugleClientOption) (FugleClient, error) {
	instance := &fugleClient{
		host:          "https://api.fugle.tw",
		headerAccept:  "*/*",
		chartEndpoint: "/realtime/v0/intraday/chart",
	}
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.config == nil {
		log.Fatal("Config object not provided")
		return nil, errors.New("config object not provided")
	}
	return instance, nil
}

func (cli *fugleClient) Chart(symbolID string, oddLot bool) FugleAPIResponse {

	url := fmt.Sprintf(
		"%s%s?apiToken=%s&symbolId=%s",
		cli.host,
		cli.chartEndpoint,
		cli.config.GetFugleConfig().GetAPIToken(),
		symbolID,
	)
	if oddLot {
		url = url + "&oddLot=true"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", cli.headerAccept)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fugleAPIResponse := FugleAPIResponse{}
	json.NewDecoder(resp.Body).Decode(&fugleAPIResponse)
	return fugleAPIResponse
}
