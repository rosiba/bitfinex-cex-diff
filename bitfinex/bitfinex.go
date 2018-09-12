package bitfinex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Ticker struct {
	Symbol          string
	Bid             float64
	BidSize         float64
	Ask             float64
	AskSize         float64
	DailyChange     float64
	DailyChangePerc float64
	LastPrice       float64
	Volume          float64
	High            float64
	Low             float64
}

func GetTickers() []Ticker {
	var foo [][]interface{}
	var client = &http.Client{Timeout: 10 * time.Second}
	data, err := client.Get("https://api.bitfinex.com/v2/tickers?symbols=tBTCUSD,tETHUSD,tBCHUSD,tBTGUSD,tDSHUSD,tXRPUSD,tZECUSD")
	if err != nil {
		fmt.Println("Unable to get ticker info.")
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	json.Unmarshal(body, &foo)

	var tickers = make([]Ticker, len(foo))

	for i, r := range foo {
		tickers[i] = Ticker{
			r[0].(string),
			r[1].(float64),
			r[2].(float64),
			r[3].(float64),
			r[4].(float64),
			r[5].(float64),
			r[6].(float64),
			r[7].(float64),
			r[8].(float64),
			r[9].(float64),
			r[10].(float64),
		}
	}
	return tickers
}
