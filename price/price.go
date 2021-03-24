package price

import (
	"stock-alerts/util"
)

type PriceApiResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    PriceData `json:"data"`
}

type PriceData struct {
	CurrentPrice  string `json:"pricecurrent"`
	Company       string `json:"company"`
	LastUpdatedAt string `json:"lastupd_epoch"`
}

func GetPrice(symbol string) []byte {

	responseData := util.Api(symbol)

	// var priceApiResponse PriceApiResponse
	// json.Unmarshal(responseData, &priceApiResponse)

	return responseData
}
