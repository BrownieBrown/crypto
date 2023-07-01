package api_test

// TODO: Fix integration tests
//
//import (
//	"crypto/pkg/api"
//	"crypto/pkg/config"
//	"crypto/pkg/constants"
//	"fmt"
//	"github.com/adshao/go-binance/v2"
//	"github.com/stretchr/testify/assert"
//	"strconv"
//	"testing"
//)
//
//func setupTestClient() *api.BinanceClient {
//	binance.UseTestnet = true
//	config.LoadEnv()
//	apiKey := config.GetEnvVar("TEST_BINANCE_API_KEY")
//	secretKey := config.GetEnvVar("TEST_BINANCE_SECRET_KEY")
//	fmt.Println("API Key: ", apiKey)
//	fmt.Println("Secret Key: ", secretKey)
//	binanceClient := api.NewBinanceClient(apiKey, secretKey)
//	return binanceClient
//}
//
//func TestGetAccountInfo(t *testing.T) {
//	client := setupTestClient()
//
//	testCases := []struct {
//		name          string
//		expectedError bool
//	}{
//		{
//			name:          "get account info",
//			expectedError: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			account, err := client.GetAccountInfo()
//
//			if tc.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, account)
//			}
//		})
//	}
//}
//
//func TestGetCurrentPrice(t *testing.T) {
//	client := setupTestClient()
//
//	testCases := []struct {
//		name          string
//		expectedError bool
//	}{
//		{
//			name:          "get current price",
//			expectedError: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			price, err := client.GetCurrentPrice(constants.TestSymbol)
//
//			if tc.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, price)
//			}
//		})
//	}
//}
//
//func TestPlaceOrder(t *testing.T) {
//	client := setupTestClient()
//
//	testCases := []struct {
//		name          string
//		expectedError bool
//	}{
//		{
//			name:          "place order",
//			expectedError: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			order, err := client.PlaceOrder(constants.TestSymbol, constants.TestSide, constants.TestQuantity)
//
//			if tc.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, order)
//				assert.Equal(t, constants.TestSymbol, order.Symbol)
//				assert.Equal(t, binance.SideType(constants.TestSide), order.Side)
//				assert.Equal(t, strconv.FormatFloat(constants.TestQuantity, 'f', 8, 64), order.OrigQuantity)
//			}
//		})
//	}
//}
//
//func TestGetOrderStatus(t *testing.T) {
//	client := setupTestClient()
//
//	testCases := []struct {
//		name          string
//		expectedError bool
//	}{
//		{
//			name:          "get order status",
//			expectedError: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			order, err := client.PlaceOrder(constants.TestSymbol, constants.TestSide, constants.TestQuantity)
//
//			if tc.expectedError {
//				assert.Error(t, err)
//				return
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, order)
//				assert.Equal(t, constants.TestSymbol, order.Symbol)
//				assert.Equal(t, binance.SideType(constants.TestSide), order.Side)
//				assert.Equal(t, strconv.FormatFloat(constants.TestQuantity, 'f', 8, 64), order.OrigQuantity)
//			}
//			orderId := order.OrderID
//			response, err := client.GetOrderStatus(constants.TestSymbol, orderId)
//
//			if tc.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, order)
//				assert.NotNil(t, response)
//				assert.Equal(t, constants.TestSymbol, order.Symbol)
//				assert.Equal(t, binance.SideType(constants.TestSide), order.Side)
//				assert.Equal(t, strconv.FormatFloat(constants.TestQuantity, 'f', 8, 64), order.OrigQuantity)
//				assert.Equal(t, constants.TestSymbol, response.Symbol)
//				assert.Equal(t, orderId, response.OrderID)
//			}
//		})
//	}
//}
//
//func TestGetTradeHistory(t *testing.T) {
//	client := setupTestClient()
//
//	testCases := []struct {
//		name          string
//		expectedError bool
//	}{
//		{
//			name:          "get order status",
//			expectedError: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			history, err := client.GetTradeHistory(constants.TestSymbol)
//
//			if tc.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, history)
//			}
//		})
//	}
//}
//
//func TestHistoricalData(t *testing.T) {
//	client := setupTestClient()
//
//	testCases := []struct {
//		name          string
//		expectedError bool
//	}{
//		{
//			name:          "get historical data",
//			expectedError: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			history, err := client.GetHistoricalData(constants.TestSymbol, constants.TestInterval, constants.TestLimit)
//
//			if tc.expectedError {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//				assert.NotNil(t, history)
//			}
//		})
//	}
//}
