package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RainrainWu/fugle-realtime-go/config"
	"github.com/RainrainWu/fugle-realtime-go/logger"
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
		host:           "api.fugle.tw",
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
		logger.PrintLogger.Fatal("Config object not provided")
		return nil, errors.New("config object not provided")
	}
	return instance, nil
}

func (cli *fugleClient) callAPI(endpoint, symbolID string, oddLot bool) *http.Response {

	targetURL, err := url.Parse(fmt.Sprintf("https://%s%s", cli.host, endpoint))
	if err != nil {
		logger.PrintLogger.Error(err.Error())
		return nil
	}
	params := url.Values{}
	params.Add("apiToken", cli.config.GetFugleConfig().GetAPIToken())
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	targetURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", targetURL.String(), nil)
	if err != nil {
		logger.PrintLogger.Error(err.Error())
		return nil
	}
	req.Header.Set("accept", cli.headerAccept)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.PrintLogger.Error(err.Error())
		return nil
	}
	return resp
}

func (cli *fugleClient) closeReponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		logger.PrintLogger.Error(err.Error())
	}
}

func (cli *fugleClient) decodeResponseBody(respBody io.ReadCloser) FugleAPIResponse {

	fugleAPIResponse := FugleAPIResponse{}
	err := json.NewDecoder(respBody).Decode(&fugleAPIResponse)
	if err != nil {
		logger.PrintLogger.Error(err.Error())
	}
	return fugleAPIResponse
}

func (cli *fugleClient) Chart(symbolID string, oddLot bool) FugleAPIResponse {

	resp := cli.callAPI(cli.chartEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(resp.Body)
	if resp.StatusCode != http.StatusOK {
		digest, _ := ioutil.ReadAll(resp.Body)
		logger.PrintLogger.Error("unexpected response: ", string(digest))
		return FugleAPIResponse{}
	}
	return cli.decodeResponseBody(resp.Body)
}

func (cli *fugleClient) Quote(symbolID string, oddLot bool) FugleAPIResponse {

	resp := cli.callAPI(cli.quoteEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(resp.Body)
	if resp.StatusCode != http.StatusOK {
		digest, _ := ioutil.ReadAll(resp.Body)
		logger.PrintLogger.Error("unexpected response: ", string(digest))
		return FugleAPIResponse{}
	}
	return cli.decodeResponseBody(resp.Body)
}

func (cli *fugleClient) Meta(symbolID string, oddLot bool) FugleAPIResponse {

	resp := cli.callAPI(cli.metaEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(resp.Body)
	if resp.StatusCode != http.StatusOK {
		digest, _ := ioutil.ReadAll(resp.Body)
		logger.PrintLogger.Error("unexpected response: ", string(digest))
		return FugleAPIResponse{}
	}
	return cli.decodeResponseBody(resp.Body)
}

func (cli *fugleClient) Dealts(symbolID string, oddLot bool) FugleAPIResponse {

	resp := cli.callAPI(cli.dealtsEndpoint, symbolID, oddLot)
	defer cli.closeReponseBody(resp.Body)
	if resp.StatusCode != http.StatusOK {
		digest, _ := ioutil.ReadAll(resp.Body)
		logger.PrintLogger.Error("unexpected response: ", string(digest))
		return FugleAPIResponse{}
	}
	return cli.decodeResponseBody(resp.Body)
}
