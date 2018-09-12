package cex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Ticker struct {
	Timestamp string  `json:"timestamp"`
	Pair      string  `json:"pair"`
	Low       float64 `json:"low,string"`
	High      float64 `json:"high,string"`
	Last      float64 `json:"last,string"`
	Volume    float64 `json:"volume,string"`
	Volume30D float64 `json:"volume30d,string"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
}

type Tickers struct {
	E    string   `json:"e"`
	Ok   string   `json:"ok"`
	Data []Ticker `json:"data"`
}

func GetTickers() Tickers {
	var tickers Tickers
	var client = &http.Client{Timeout: 10 * time.Second}
	data, err := client.Get("https://cex.io/api/tickers/USD")
	if err != nil {
		fmt.Println("Unable to get ticker information.")
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	json.Unmarshal(body, &tickers)
	return tickers
}
