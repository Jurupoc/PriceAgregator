package domain

import "time"

type PriceData struct {
	Name         string
	Symbol       string
	UsdPrice     float64
	UsdMarketCap float64
	RetrieveAt   time.Time
}
