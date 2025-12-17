package domain

import "time"

// PriceSnapshot representa um snapshot normalizado de preço de criptomoeda
type PriceSnapshot struct {
	Symbol    string
	PriceUSD  float64
	Source    string
	Timestamp time.Time
}

// PriceData mantido para compatibilidade durante refatoração
type PriceData struct {
	Name         string
	Symbol       string
	UsdPrice     float64
	UsdMarketCap float64
	RetrieveAt   time.Time
}
