package client

import (
	"encoding/json"
	"fmt"
	"log"
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
	Date          string `json:"date"`
	Mode          string `json:"mode"`
	SymbolID      string `json:"symbolId"`
	CountryCode   string `json:"countryCode"`
	Timezone      string `json:"timeZone"`
	LastUpdatedAt string `json:"lastUpdatedAt"`
}

type FugleAPIChartPrice struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int     `json:"volume"`
	Unit   int     `json:"unit"`
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
	At     string `json:"at"`
	Unit   int    `json:"unit"`
	Volume int    `json:"volume"`
}

type FugleAPITrial struct {
	At     string  `json:"at"`
	Price  float64 `json:"price"`
	Unit   int     `json:"unit"`
	Volume int     `json:"volume"`
}

type FugleAPITrade struct {
	At     string  `json:"at"`
	Price  float64 `json:"price"`
	Unit   int     `json:"unit"`
	Volume int     `json:"volume"`
	Serial int     `json:"serial"`
}

type FugleAPIOrder struct {
	At       string              `json:"at"`
	Bestbids []FugleAPIBestPrice `json:"bestBids"`
	Bestasks []FugleAPIBestPrice `json:"bestAsks"`
}

type FugleAPIBestPrice struct {
	Price  float64 `json:"price"`
	Unit   int     `json:"unit"`
	Volume int     `json:"volume"`
}

type FugleAPIQuotePrice struct {
	Price float64 `json:"price"`
	At    string  `json:"at"`
}

type FugleAPIMeta struct {
	Isindex                bool   `json:"isIndex"`
	Namezhtw               string `json:"nameZhTw"`
	Industryzhtw           string `json:"industryZhTw"`
	Pricereference         int    `json:"priceReference"`
	Pricehighlimit         int    `json:"priceHighLimit"`
	Pricelowlimit          int    `json:"priceLowLimit"`
	Candaybuysell          bool   `json:"canDayBuySell"`
	Candaysellbuy          bool   `json:"canDaySellBuy"`
	Canshortmargin         bool   `json:"canShortMargin"`
	Canshortlend           bool   `json:"canShortLend"`
	Volumeperunit          int    `json:"volumePerUnit"`
	Currency               string `json:"currency"`
	Isterminated           bool   `json:"isTerminated"`
	Issuspended            bool   `json:"isSuspended"`
	Iswarrant              bool   `json:"isWarrant"`
	Typezhtw               string `json:"typeZhTw"`
	Abnormal               string `json:"abnormal"`
	IsUnusuallyRecommended bool   `json:"isUnusuallyRecommended"`
}

type FugleAPIDealts struct {
	At     string  `json:"at"`
	Price  float64 `json:"price"`
	Unit   int     `json:"unit"`
	Serial int     `json:"serial"`
}

func (resp *FugleAPIResponse) PrettyPrint() {
	buffer, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(buffer))
}
