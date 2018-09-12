package main

import (
	"fmt"
	"github.com/rosiba/bitfinex-cex-diff/bitfinex"
	"github.com/rosiba/bitfinex-cex-diff/cex"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Coin struct {
	Name     string
	CexPrice float64
	BfxPrice float64
	Stats    CoinStats
}

type CoinStats struct {
	Ratio      float64
	MinRatio   float64
	MaxRatio   float64
	AvgRatio   float64
	SumRatio   float64
	CountRatio float64
}

func PercentDiff(a float64, b float64) float64 {
	return (a/b - 1) * 100
}

func InitializeCoinData() []Coin {
	var coins []Coin
	coinList := []string{"BTC", "ETH", "BCH", "BTG", "DASH", "XRP", "ZEC"}
	bfxTickers := bitfinex.GetTickers()
	cexTickers := cex.GetTickers()
	for i := 0; i < len(coinList); i++ {
		initialRatio := PercentDiff(cexTickers.Data[i].Last, bfxTickers[i].LastPrice)
		coin := Coin{
			Name:     coinList[i],
			CexPrice: cexTickers.Data[i].Last,
			BfxPrice: bfxTickers[i].LastPrice,
			Stats: CoinStats{
				Ratio:      initialRatio,
				MinRatio:   initialRatio,
				MaxRatio:   initialRatio,
				SumRatio:   initialRatio,
				AvgRatio:   initialRatio,
				CountRatio: 1,
			},
		}
		coins = append(coins, coin)
	}
	return coins
}

func UpdateCoinData(coins []Coin) {
	coinList := []string{"BTC", "ETH", "BCH", "BTG", "DASH", "XRP", "ZEC"}
	bfxTickers := bitfinex.GetTickers()
	cexTickers := cex.GetTickers()
	var c *Coin
	for i := 0; i < len(coinList); i++ {
		c = &coins[i]
		c.BfxPrice = bfxTickers[i].LastPrice
		c.CexPrice = cexTickers.Data[i].Last
		c.Stats.Ratio = PercentDiff(c.CexPrice, c.BfxPrice)
		if c.Stats.Ratio < c.Stats.MinRatio {
			c.Stats.MinRatio = c.Stats.Ratio
		}
		if c.Stats.Ratio > c.Stats.MaxRatio {
			c.Stats.MaxRatio = c.Stats.Ratio
		}
		c.Stats.SumRatio += c.Stats.Ratio
		c.Stats.CountRatio++
		c.Stats.AvgRatio = c.Stats.SumRatio / c.Stats.CountRatio
	}
}

func UpdateCoins(coins []Coin) {
	for {
		time.Sleep(6 * time.Second)
		UpdateCoinData(coins)
	}
}

func PrintHeader() {
	h := []string{"CURRENCY", "CEX", "BFX", "RATIO", "AVG", "MIN", "MAX"}
	format := "|%-9s|%-9s|%-9s|%-9s|%-9s|%-9s|%-9s|\n"
	fmt.Printf(format, h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7])
}

func PrintCoin(coin Coin) {
	c := coin
	format := "|%-9s|%-9.2f|%-9.2f|%%%-8.2f|%%%-8.2f|%%%-8.2f|%%%-8.2f|\n"
	fmt.Printf(format, c.Name, c.CexPrice, c.BfxPrice, c.Stats.Ratio, c.Stats.AvgRatio, c.Stats.MinRatio, c.Stats.MaxRatio)
}

func PrintCoins(coins []Coin) {
	for _, coin := range coins {
		PrintCoin(coin)
	}
}

func clearWindows() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearLinux() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func clear() {
	if runtime.GOOS == "windows" {
		clearWindows()
	} else if runtime.GOOS == "linux" {
		clearLinux()
	} else {
		fmt.Println("Put the calculator on your hand to the ground, slowly boi.")
		fmt.Println("Or simply install CalculOS on it.")
	}
}

func main() {
	coins := InitializeCoinData()
	go UpdateCoins(coins)
	for {
		clear()
		PrintHeader()
		PrintCoins(coins)
		time.Sleep(time.Second / 10)
	}
}
