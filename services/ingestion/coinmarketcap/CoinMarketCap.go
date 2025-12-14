package coinmarketcap

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/samber/lo"

	"github.com/Jurupoc/PriceAgregator/ingestion/internal/config"
	"github.com/Jurupoc/PriceAgregator/ingestion/internal/domain"
)

type coinMarketCapProvider struct {
	config config.Config
	client http.Client
}

func NewCoinMarketCapProvider(cf config.ConfigProvider, cl http.Client) domain.DataProvider {
	config := cf.Load()
	return &coinMarketCapProvider{
		config: config,
		client: cl,
	}
}

func (i *coinMarketCapProvider) FetchPrice() (*domain.PriceData, error) {
	var result *domain.PriceData

	req, err := http.NewRequest(http.MethodGet, COIN_MARKET_CAP_API_URL+COIN_MARKETCAP_API_ENDPOINT, nil)
	if err != nil {
		log.Fatal("Error creating request", err)
		return result, err
	}
	req.Header.Add(COIN_MARKET_CAP_API_KEY_HEADER, i.config.CoinMarketCapAPIKey)

	q := req.URL.Query()
	q.Add("limit", "5")
	req.URL.RawQuery = q.Encode()

	resp, err := i.client.Do(req)
	if err != nil {
		log.Fatal("Error sending request", err)
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
		return result, err
	}

	var response PriceDataResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal("Error unmarshalling response body:", err)
		return result, err
	}
	if response.Status.ErrorCode != nil {
		log.Fatalf("CoinMarketCap API returned error code: %d, message: %s",
			*response.Status.ErrorCode,
			lo.FromPtr(response.Status.ErrorMessage),
		)
	}
	log.Printf("CoinMarketCap response: %+v \n", response)

	result = converToPriceData(response.Data[0])

	return result, nil
}

func converToPriceData(coinData CryptoAsset) *domain.PriceData {
	return &domain.PriceData{
		Name:       coinData.Name,
		Symbol:     coinData.Symbol,
		RetrieveAt: time.Now(),
		UsdPrice:   coinData.Quote["USD"].Price,
	}
}
