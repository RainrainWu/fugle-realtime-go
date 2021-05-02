package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RainrainWu/fugle-realtime-go/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type FugleClient interface {
	Chart(symbolID string, oddLot bool) (FugleAPIResponse, error)
	Quote(symbolID string, oddLot bool) (FugleAPIResponse, error)
	Meta(symbolID string, oddLot bool) (FugleAPIResponse, error)
	Dealts(symbolID string, oddLot bool) (FugleAPIResponse, error)
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
		instance.config = config.Config
	}
	return instance, nil
}

func (cli *fugleClient) callAPI(endpoint, symbolID string, oddLot bool) (*http.Response, error) {

	targetURL, err := url.Parse(fmt.Sprintf("https://%s%s", cli.host, endpoint))
	if err != nil {
		logrus.Error(err.Error())
		return nil, errors.Wrap(err, "parse url failed")
	}
	params := url.Values{}
	params.Add("apiToken", cli.config.GetFugleConfig().GetAPIToken())
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	targetURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", targetURL.String(), nil)
	if err != nil {
		logrus.Error(err.Error())
		return nil, errors.Wrap(err, "initialize request failed")
	}
	req.Header.Set("accept", cli.headerAccept)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Error(err.Error())
		return nil, errors.Wrap(err, "send request failed")
	}
	return resp, nil
}

func (cli *fugleClient) closeReponseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func (cli *fugleClient) decodeResponseBody(resp *http.Response) (FugleAPIResponse, error) {

	fugleAPIResponse := FugleAPIResponse{
		StatusCode: resp.StatusCode,
	}
	defer cli.closeReponseBody(resp.Body)
	if resp.StatusCode != http.StatusOK {
		digest, _ := ioutil.ReadAll(resp.Body)
		logrus.WithField("status code", resp.StatusCode).Error("unexpected response: ", string(digest))
		return fugleAPIResponse, errors.New("unexpected status code")
	}

	err := json.NewDecoder(resp.Body).Decode(&fugleAPIResponse)
	if err != nil {
		logrus.Error(err.Error())
		return fugleAPIResponse, errors.Wrap(err, "decode json failed")
	}
	return fugleAPIResponse, nil
}

func (cli *fugleClient) Chart(symbolID string, oddLot bool) (FugleAPIResponse, error) {

	resp, err := cli.callAPI(cli.chartEndpoint, symbolID, oddLot)
	if err != nil {
		return FugleAPIResponse{}, errors.Wrap(err, "call api failed")
	}
	return cli.decodeResponseBody(resp)
}

func (cli *fugleClient) Quote(symbolID string, oddLot bool) (FugleAPIResponse, error) {

	resp, err := cli.callAPI(cli.quoteEndpoint, symbolID, oddLot)
	if err != nil {
		return FugleAPIResponse{}, errors.Wrap(err, "call api failed")
	}
	return cli.decodeResponseBody(resp)
}

func (cli *fugleClient) Meta(symbolID string, oddLot bool) (FugleAPIResponse, error) {

	resp, err := cli.callAPI(cli.metaEndpoint, symbolID, oddLot)
	if err != nil {
		return FugleAPIResponse{}, errors.Wrap(err, "call api failed")
	}
	return cli.decodeResponseBody(resp)
}

func (cli *fugleClient) Dealts(symbolID string, oddLot bool) (FugleAPIResponse, error) {

	resp, err := cli.callAPI(cli.dealtsEndpoint, symbolID, oddLot)
	if err != nil {
		return FugleAPIResponse{}, errors.Wrap(err, "call api failed")
	}
	return cli.decodeResponseBody(resp)
}
