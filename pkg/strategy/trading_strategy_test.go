package strategy_test

import (
	"crypto/pkg/mocks"
	"crypto/pkg/strategy"
	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestExecuteStrategy(t *testing.T) {
	// Instantiate the mocks
	mockAPI := new(mocks.MockBinanceClient)
	mockIndicators := new(mocks.MockIndicators)

	// Define the behavior of the GetHistoricalData mock method
	klines := []*binance.Kline{
		// Populate with test data...
	}
	mockAPI.On("GetHistoricalData", "BTCUSD", "1h", 100).Return(klines, nil)

	// Define the behavior of the mock methods
	mockIndicators.On("ExponentialMovingAverage", mock.Anything, 14).Return([]float64{1.0, 2.0, 3.0})
	// Define similar behavior for the other indicator methods...
	// ...

	mockAPI.On("PlaceOrder", "BTCUSD", "BUY", mock.Anything).Return(&binance.CreateOrderResponse{}, nil)
	mockAPI.On("PlaceOrder", "BTCUSD", "SELL", mock.Anything).Return(&binance.CreateOrderResponse{}, nil)

	// Create the TradingStrategy
	ts := strategy.NewTradingStrategy(mockAPI, mockIndicators, "BTCUSD", "1h", 14, 100, 0.02, 0.06)

	// Call the method we're testing
	ts.ExecuteStrategy()

	// Assert that the mock was called as expected
	mockAPI.AssertExpectations(t)
	mockIndicators.AssertExpectations(t)
}
