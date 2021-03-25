package price

import (
	"stock-alerts/util"
)

func GetPrice(symbol string) []byte {
	responseData := util.Api(symbol)
	return responseData
}
