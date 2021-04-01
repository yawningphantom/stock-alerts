package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stock-alerts/price"
	"stock-alerts/util"
)

type PriceHandler struct {
	l *log.Logger
}

func NewPriceHandler(l *log.Logger) *PriceHandler {
	return &PriceHandler{l}
}

type PriceHandlerInput struct {
	Symbol   string        `json:"code"`
	Exchange util.Exchange `json:"exchange"`
}

func (h *PriceHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Price Handler Called")
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Price Handler Failed !!!", http.StatusBadRequest)
		return
	}

	var priceHandlerInput PriceHandlerInput
	json.Unmarshal(input, &priceHandlerInput)

	var Symbol string = priceHandlerInput.Symbol
	var Exchange util.Exchange = priceHandlerInput.Exchange

	stockUrl, _ := price.GenerateStockUrl(Symbol, Exchange)
	fmt.Println(stockUrl)
	stockData, _ := price.GetStockData(stockUrl)
	if err != nil {
		http.Error(rw, "Price Handler Failed !!!", http.StatusBadRequest)
		return
	}

	apiResponse, err := json.Marshal(stockData)
	if err != nil {
		fmt.Println(err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(apiResponse)
}
