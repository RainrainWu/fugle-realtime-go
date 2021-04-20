package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RainrainWu/fugle-realtime-go/config"
)

type FugleClient interface {
	Chart(symbolID string, oddLot bool) FugleAPIResponse
	Quote(symbolID string, oddLot bool) FugleAPIResponse
	Meta(symbolID string, oddLot bool) FugleAPIResponse
	Dealts(symbolID string, oddLot bool) FugleAPIResponse
}

type fugleClient struct {
	host           string
	headerAccept   string
	chartEndpoint  string
	quoteEndpoint  string
	metaEndpoint   string
	dealtsEndpoint string
	config         config.ConfigSet
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
		host:           "https://api.fugle.tw",
		headerAccept:   "*/*",
		chartEndpoint:  "/realtime/v0/intraday/chart",
		quoteEndpoint:  "/realtime/v0/intraday/quote",
		metaEndpoint:   "/realtime/v0/intraday/meta",
		dealtsEndpoint: "/realtime/v0/intraday/dealts",
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

func (cli *fugleClient) callAPI(endpoint, symbolID string, oddLot bool) io.ReadCloser {

	targetURL, err := url.Parse(fmt.Sprintf("%s%s", cli.host, endpoint))
	if err != nil {
		log.Fatal(err.Error())
	}
	params := url.Values{}
	params.Add("apiToken", cli.config.GetFugleConfig().GetAPIToken())
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	targetURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", targetURL.String(), nil)
	fmt.Println(req.URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("accept", cli.headerAccept)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	return resp.Body
}

func (cli *fugleClient) closeReponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (cli *fugleClient) decodeResponseBody(respBody io.ReadCloser) FugleAPIResponse {

	fugleAPIResponse := FugleAPIResponse{}
	err := json.NewDecoder(respBody).Decode(&fugleAPIResponse)
	if err != nil {
		log.Fatal(err.Error())
	}
	return fugleAPIResponse
}

func (cli *fugleClient) Chart(symbolID string, oddLot bool) FugleAPIResponse {

	respBody := cli.callAPI(cli.chartEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(respBody)
	return cli.decodeResponseBody(respBody)
}

func (cli *fugleClient) Quote(symbolID string, oddLot bool) FugleAPIResponse {

	respBody := cli.callAPI(cli.quoteEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(respBody)
	return cli.decodeResponseBody(respBody)
}

func (cli *fugleClient) Meta(symbolID string, oddLot bool) FugleAPIResponse {

	respBody := cli.callAPI(cli.metaEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(respBody)
	return cli.decodeResponseBody(respBody)
}

func (cli *fugleClient) Dealts(symbolID string, oddLot bool) FugleAPIResponse {

	respBody := cli.callAPI(cli.dealtsEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(respBody)
	return cli.decodeResponseBody(respBody)
}
