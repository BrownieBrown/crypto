package api_test

import (
	"crypto/pkg/constants"
	"crypto/pkg/mocks"
	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMockBinanceClient(t *testing.T) {
	mockClient := mocks.SetupMockBinanceClient()

	testCases := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{
			name: "TestGetAccountInfo_Success",
			testFunc: func(t *testing.T) {
				account, err := mockClient.GetAccountInfo()
				assert.NoError(t, err)
				assert.Equal(t, constants.TestMakerComm, account.MakerCommission)
				assert.Equal(t, constants.TestMakerComm, account.TakerCommission)
				assert.Equal(t, int64(0), account.BuyerCommission)
				assert.Equal(t, "SPOT", account.AccountType)
				assert.Equal(t, constants.TestAsset, account.Balances[0].Asset)
				assert.Equal(t, constants.TestFree, account.Balances[0].Free)
			},
		},
		{
			name: "TestGetAccountInfo_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.GetAccountInfoError()

				assert.Error(t, err)
				assert.Equal(t, "failed to get account info", err.Error())
			},
		},
		{
			name: "TestGetCurrentPrice_Success",
			testFunc: func(t *testing.T) {
				price, err := mockClient.GetCurrentPrice(constants.TestSymbol)
				assert.NoError(t, err)
				assert.Equal(t, 50000.0, price)
			},
		},
		{
			name: "TestGetCurrentPrice_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.GetCurrentPriceError(constants.TestSymbol)
				assert.Error(t, err)
				assert.Equal(t, "failed to get current price", err.Error())
			},
		},
		{
			name: "TestPlaceOrder_Success",
			testFunc: func(t *testing.T) {
				order, err := mockClient.PlaceOrder(constants.TestSymbol, constants.TestSide, constants.TestQuantity)
				assert.NoError(t, err)
				assert.Equal(t, constants.TestSymbol, order.Symbol)
				assert.Equal(t, binance.SideType(constants.TestSide), order.Side)
				assert.Equal(t, strconv.FormatFloat(constants.TestQuantity, 'f', 6, 64), order.OrigQuantity)
			},
		},
		{
			name: "TestPlaceOrder_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.PlaceOrderError(constants.TestSymbol, constants.TestSide, constants.TestQuantity)
				assert.Error(t, err)
				assert.Equal(t, "failed to place order", err.Error())
			},
		},
		{
			name: "TestCancelOrder_Success",
			testFunc: func(t *testing.T) {
				order, err := mockClient.CancelOrder(constants.TestSymbol, constants.TestOrderID)
				assert.NoError(t, err)
				assert.Equal(t, constants.TestSymbol, order.Symbol)
				assert.Equal(t, constants.TestOrderID, order.OrderID)
			},
		},
		{
			name: "TestCancelOrder_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.CancelOrderError(constants.TestSymbol, constants.TestOrderID)
				assert.Error(t, err)
				assert.Equal(t, "failed to cancel order", err.Error())
			},
		},
		{
			name: "TestGetOrderStatus_Success",
			testFunc: func(t *testing.T) {
				order, err := mockClient.GetOrderStatus(constants.TestSymbol, constants.TestOrderID)
				assert.NoError(t, err)
				assert.Equal(t, constants.TestSymbol, order.Symbol)
				assert.Equal(t, constants.TestOrderID, order.OrderID)
			},
		},
		{
			name: "TestGetOrderStatus_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.GetOrderStatusError(constants.TestSymbol, constants.TestOrderID)
				assert.Error(t, err)
				assert.Equal(t, "failed to get order status", err.Error())
			},
		},
		{
			name: "TestListOpenOrders_Success",
			testFunc: func(t *testing.T) {
				orders, err := mockClient.ListOpenOrders(constants.TestSymbol)
				assert.NoError(t, err)
				assert.Equal(t, constants.TestSymbol, orders[0].Symbol)
				assert.Equal(t, len(orders), 2)
			},
		},
		{
			name: "TestListOpenOrders_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.ListOpenOrdersError(constants.TestSymbol)
				assert.Error(t, err)
				assert.Equal(t, "failed to get open orders", err.Error())
			},
		},
		{
			name: "TestListOrders_Success",
			testFunc: func(t *testing.T) {
				orders, err := mockClient.ListOrders(constants.TestSymbol)
				assert.NoError(t, err)
				assert.Equal(t, constants.TestSymbol, orders[0].Symbol)
				assert.Equal(t, len(orders), 2)
			},
		},
		{
			name: "TestListOrders_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.ListOrdersError(constants.TestSymbol)
				assert.Error(t, err)
				assert.Equal(t, "failed to get orders", err.Error())
			},
		},
		{
			name: "TestGetTradeHistory_Success",
			testFunc: func(t *testing.T) {
				trades, err := mockClient.GetTradeHistory(constants.TestSymbol)
				assert.NoError(t, err)
				assert.Equal(t, constants.TestSymbol, trades[0].Symbol)
				assert.Equal(t, len(trades), 2)
			},
		},
		{
			name: "TestGetTradeHistory_Failure",
			testFunc: func(t *testing.T) {
				_, err := mockClient.GetTradeHistoryError(constants.TestSymbol)
				assert.Error(t, err)
				assert.Equal(t, "failed to get trade history", err.Error())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testFunc)
	}
}
