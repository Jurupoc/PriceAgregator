package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jurupoc/PriceAgregator/ingestion/coinmarketcap"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/config"
)

func main() {
	var httpClient http.Client

	cf := config.NewConfigProvider()

	coinCapProvider := coinmarketcap.NewCoinMarketCapProvider(cf, httpClient)
	prices, err := coinCapProvider.FetchPrice()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prices)
}
