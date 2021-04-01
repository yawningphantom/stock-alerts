package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"stock-alerts/price"

	"github.com/robfig/cron/v3"
)

func InitialiseCron() {
	c := cron.New()
	c.AddFunc("CRON_TZ=Asia/Kolkata @every 1m", GetPricesForStocksAtAlert) // TODO : improve the cron timing to run only when stock market is open
	c.Start()
}

func GetPricesForStocksAtAlert() {
	plan, _ := ioutil.ReadFile("stocksAtAlert.json")
	var data price.StocksAtAlert
	json.Unmarshal(plan, &data)

	for _, stock := range data.Stocks {
		fmt.Println("Checking for stock :: ", stock.Symbol, " :: ", stock)
		go price.StockChecker(stock)
	}
}
