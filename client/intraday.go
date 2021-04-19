package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	Info   FugleAPIInfo                  `json:"info"`
	Chart  map[string]FugleAPIChartPrice `json:"chart"`
	Quote  FugleAPIQuote                 `json:"quote"`
	Meta   FugleAPIMeta                  `json:"meta"`
	Dealts []FugleAPIDealts              `json:"dealts"`
}

type FugleAPIInfo struct {
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	Date          string    `json:"date"`
	Mode          string    `json:"mode"`
	SymbolID      string    `json:"symbolId"`
	CountryCode   string    `json:"countryCode"`
	Timezone      string    `json:"timeZone"`
}

type FugleAPIChartPrice struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Unit   int     `json:"unit"`
	Volume int     `json:"volume"`
}

type FugleAPIQuote struct {
	Iscurbing      bool               `json:"isCurbing"`
	Iscurbingrise  bool               `json:"isCurbingRise"`
	Iscurbingfall  bool               `json:"isCurbingFall"`
	Istrial        bool               `json:"isTrial"`
	Isopendelayed  bool               `json:"isOpenDelayed"`
	Isclosedelayed bool               `json:"isCloseDelayed"`
	Ishalting      bool               `json:"isHalting"`
	Isclosed       bool               `json:"isClosed"`
	Total          FugleAPITotal      `json:"total"`
	Trial          FugleAPITrial      `json:"trial"`
	Trade          FugleAPITrade      `json:"trade"`
	Order          FugleAPIOrder      `json:"order"`
	PriceHigh      FugleAPIQuotePrice `json:"priceHigh"`
	PriceLow       FugleAPIQuotePrice `json:"priceLow"`
	PriceOpen      FugleAPIQuotePrice `json:"priceOpen"`
}

type FugleAPITotal struct {
	At     time.Time `json:"at"`
	Order  int       `json:"order"`
	Price  int       `json:"price"`
	Unit   int       `json:"unit"`
	Volume int       `json:"volume"`
}

type FugleAPITrial struct {
	At     time.Time `json:"at"`
	Price  int       `json:"price"`
	Unit   int       `json:"unit"`
	Volume int       `json:"volume"`
}

type FugleAPITrade struct {
	At     time.Time `json:"at"`
	Price  int       `json:"price"`
	Unit   int       `json:"unit"`
	Volume int       `json:"volume"`
	Serial int       `json:"serial"`
}

type FugleAPIOrder struct {
	At       time.Time           `json:"at"`
	Bestbids []FugleAPIBestPrice `json:"bestBids"`
	Bestasks []FugleAPIBestPrice `json:"bestAsks"`
}

type FugleAPIBestPrice struct {
	Price  int `json:"price"`
	Unit   int `json:"unit"`
	Volume int `json:"volume"`
}

type FugleAPIQuotePrice struct {
	Price int       `json:"price"`
	At    time.Time `json:"at"`
}

type FugleAPIMeta struct {
	Isindex        bool   `json:"isIndex"`
	Namezhtw       string `json:"nameZhTw"`
	Industryzhtw   string `json:"industryZhTw"`
	Pricereference int    `json:"priceReference"`
	Pricehighlimit int    `json:"priceHighLimit"`
	Pricelowlimit  int    `json:"priceLowLimit"`
	Candaybuysell  bool   `json:"canDayBuySell"`
	Candaysellbuy  bool   `json:"canDaySellBuy"`
	Canshortmargin bool   `json:"canShortMargin"`
	Canshortlend   bool   `json:"canShortLend"`
	Volumeperunit  int    `json:"volumePerUnit"`
	Currency       string `json:"currency"`
	Isterminated   bool   `json:"isTerminated"`
	Issuspended    bool   `json:"isSuspended"`
	Iswarrant      bool   `json:"isWarrant"`
	Typezhtw       string `json:"typeZhTw"`
	Abnormal       string `json:"abnormal"`
}

type FugleAPIDealts struct {
	At     time.Time `json:"at"`
	Price  float64   `json:"price"`
	Unit   int       `json:"unit"`
	Serial int       `json:"serial"`
}

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

func (cli *fugleClient) concatURL(endpoint, symbolID string, oddLot bool) string {

	url := fmt.Sprintf(
		"%s%s?apiToken=%s&symbolId=%s",
		cli.host,
		endpoint,
		cli.config.GetFugleConfig().GetAPIToken(),
		symbolID,
	)
	if oddLot {
		url = url + "&oddLot=true"
	}
	return url
}

func (cli *fugleClient) callAPI(url string) io.ReadCloser {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", cli.headerAccept)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}

func (cli *fugleClient) Chart(symbolID string, oddLot bool) FugleAPIResponse {

	url := cli.concatURL(cli.chartEndpoint, symbolID, oddLot)
	respBody := cli.callAPI(url)
	defer respBody.Close()

	fugleAPIResponse := FugleAPIResponse{}
	json.NewDecoder(respBody).Decode(&fugleAPIResponse)
	return fugleAPIResponse
}

func (cli *fugleClient) Quote(symbolID string, oddLot bool) FugleAPIResponse {

	url := cli.concatURL(cli.quoteEndpoint, symbolID, oddLot)
	respBody := cli.callAPI(url)
	defer respBody.Close()

	fugleAPIResponse := FugleAPIResponse{}
	json.NewDecoder(respBody).Decode(&fugleAPIResponse)
	return fugleAPIResponse
}

func (cli *fugleClient) Meta(symbolID string, oddLot bool) FugleAPIResponse {

	url := cli.concatURL(cli.metaEndpoint, symbolID, oddLot)
	respBody := cli.callAPI(url)
	defer respBody.Close()

	fugleAPIResponse := FugleAPIResponse{}
	json.NewDecoder(respBody).Decode(&fugleAPIResponse)
	return fugleAPIResponse
}

func (cli *fugleClient) Dealts(symbolID string, oddLot bool) FugleAPIResponse {

	url := cli.concatURL(cli.dealtsEndpoint, symbolID, oddLot)
	respBody := cli.callAPI(url)
	defer respBody.Close()

	fugleAPIResponse := FugleAPIResponse{}
	json.NewDecoder(respBody).Decode(&fugleAPIResponse)
	return fugleAPIResponse
}
