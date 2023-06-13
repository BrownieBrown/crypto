package bot

import "crypto/pkg/api"

type TradingBot struct {
	BinanceClient *api.BinanceClient
}

func NewTradingBot(apiKey string, secretKey string) *TradingBot {
	binanceClient := api.NewBinanceClient(apiKey, secretKey)

	return &TradingBot{
		BinanceClient: binanceClient,
	}
}

func (tb *TradingBot) GetPrice(symbol string) (float64, error) {
	return tb.BinanceClient.GetCurrentPrice(symbol)
}
