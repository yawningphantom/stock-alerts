package util

const BaseUrl = "https://priceapi.moneycontrol.com"

type Exchange string

const (
	BSE Exchange = "bse"
	NSE Exchange = "nse"
)

var SymbolMap = map[string]string{
	"ABFRL": "PFR",
}
