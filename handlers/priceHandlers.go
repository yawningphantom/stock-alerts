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
	var Exchange util.Exchange
	Exchange = priceHandlerInput.Exchange

	StockPriceUrl := fmt.Sprint(util.BaseUrl, "/pricefeed/", Exchange, "/equitycash/", Symbol)
	fmt.Println(StockPriceUrl)
	data := price.GetPrice(StockPriceUrl)
	if err != nil {
		http.Error(rw, "Price Handler Failed !!!", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)

}
