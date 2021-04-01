package price

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"stock-alerts/util"
	"strconv"
)

type StocksAtAlert struct {
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	Symbol     string        `json:"symbol"`
	LowerLimit float32       `json:"lowerLimit"`
	UpperLimit float32       `json:"upperLimit"`
	Exchange   util.Exchange `json:"exchange"`
}

type GetStockApiResponse struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	StockData StockDataResponse `json:"data"`
}

type StockDataResponse struct {
	Price           string        `json:"pricecurrent"`
	PriceChange     string        `json:"pricechange"`
	Company         string        `json:"company"`
	FullCompanyName string        `json:"SC_FULLNM"`
	Exchange        util.Exchange `json:"exchange"`
}

func GetStockData(stockUrl string) (StockData StockDataResponse, err string) {
	responseData := util.Api(stockUrl)
	var stockApiResponse GetStockApiResponse
	json.Unmarshal(responseData, &stockApiResponse)
	if stockApiResponse.Code != http.StatusOK {
		err = "Stock API failed"
	}
	StockData = stockApiResponse.StockData
	StockData.Exchange = util.GetExchange(string(StockData.Exchange))

	return
}

func CheckIfStockPriceConstraintBreached(stockPrice float32, lowerLimit float32, upperLimit float32) bool {

	if stockPrice >= upperLimit {
		fmt.Println("Stock Price upperLimit Constraint breached")
		return true
	}

	if stockPrice <= lowerLimit {
		fmt.Println("Stock Price lowerLimit Constraint breached")
		return true
	}

	return false
}

func GenerateStockUrl(Symbol string, Exchange util.Exchange) (stockUrl string, err error) {
	if Exchange != util.BSE && Exchange != util.NSE {
		Exchange = util.BSE
	}

	stockSymbol, ok := util.SymbolMap[Symbol]
	if !ok {
		return stockUrl, errors.New("invalid Stock Symbol")
	}

	stockUrl = fmt.Sprint(util.BaseUrl, "/pricefeed/", Exchange, "/equitycash/", stockSymbol)
	return
}

func StockChecker(stock Stock) {
	stockPrice := GetStockPrice(stock.Symbol, stock.Exchange)
	CheckIfStockPriceConstraintBreached(stockPrice, stock.LowerLimit, stock.UpperLimit)
}

func GetStockPrice(Symbol string, Exchange util.Exchange) float32 {
	stockUrl, _ := GenerateStockUrl(Symbol, Exchange)
	stockData, _ := GetStockData(stockUrl)
	stockPrice, _ := strconv.ParseFloat(stockData.Price, 32)
	return float32(stockPrice)
}
