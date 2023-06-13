package mocks

import (
	"crypto/pkg/api"
	"crypto/pkg/constants"
	"errors"
	"github.com/adshao/go-binance/v2"
	"strconv"
)

type MockBinanceClient struct {
	api.BinanceAPI
}

func SetupMockBinanceClient() *MockBinanceClient {
	return &MockBinanceClient{}
}

func (m *MockBinanceClient) GetAccountInfo() (*binance.Account, error) {
	account := &binance.Account{
		MakerCommission:  constants.TestMakerComm,
		TakerCommission:  constants.TestMakerComm,
		BuyerCommission:  0,
		SellerCommission: 0,
		CanTrade:         true,
		CanWithdraw:      false,
		CanDeposit:       false,
		UpdateTime:       123456789,
		AccountType:      "SPOT",
		Balances: []binance.Balance{
			{
				Asset:  constants.TestAsset,
				Free:   constants.TestFree,
				Locked: "0.00000000",
			},
		},
	}
	return account, nil
}

func (m *MockBinanceClient) GetAccountInfoError() (*binance.Account, error) {
	return nil, errors.New("failed to get account info")
}

func (m *MockBinanceClient) GetCurrentPrice(symbol string) (float64, error) {
	return 50000.0, nil
}

func (m *MockBinanceClient) GetCurrentPriceError(symbol string) (float64, error) {
	return 0, errors.New("failed to get current price")
}

func (m *MockBinanceClient) PlaceOrder(symbol string, side string, quantity float64) (*binance.CreateOrderResponse, error) {
	// Use parameters in method body
	order := &binance.CreateOrderResponse{
		Symbol:                   symbol,
		OrderID:                  10,
		ClientOrderID:            "wOceeeOzNORyLiQfw7jd8S",
		TransactTime:             1507725176595,
		Price:                    "0.02000000",
		OrigQuantity:             strconv.FormatFloat(quantity, 'f', 6, 64),
		ExecutedQuantity:         "0.000000002",
		CummulativeQuoteQuantity: "0.000000002",
		IsIsolated:               false,
		Status:                   "NEW",
		TimeInForce:              "GTC",
		Type:                     "LIMIT",
		Side:                     binance.SideType(side),
	}
	return order, nil
}

func (m *MockBinanceClient) PlaceOrderError(symbol string, side string, quantity float64) (*binance.CreateOrderResponse, error) {
	// A method to simulate an error
	return nil, errors.New("failed to place order")
}

func (m *MockBinanceClient) CancelOrder(symbol string, orderId int64) (*binance.CancelOrderResponse, error) {
	response := &binance.CancelOrderResponse{
		Symbol:                   symbol,
		OrigClientOrderID:        "1",
		OrderID:                  orderId,
		OrderListID:              0,
		ClientOrderID:            "1234",
		TransactTime:             0,
		Price:                    "100",
		OrigQuantity:             "1",
		ExecutedQuantity:         "1",
		CummulativeQuoteQuantity: "1",
		Status:                   "NEW",
		TimeInForce:              "1234",
		Type:                     "test",
		Side:                     "BUY",
	}

	return response, nil
}

func (m *MockBinanceClient) CancelOrderError(symbol string, orderId int64) (*binance.CancelOrderResponse, error) {
	return nil, errors.New("failed to cancel order")
}

func (m *MockBinanceClient) GetOrderStatus(symbol string, orderId int64) (*binance.Order, error) {
	order := &binance.Order{
		Symbol:                   symbol,
		OrderID:                  orderId,
		OrderListId:              0,
		ClientOrderID:            "",
		Price:                    "",
		OrigQuantity:             "",
		ExecutedQuantity:         "",
		CummulativeQuoteQuantity: "",
		Status:                   "",
		TimeInForce:              "",
		Type:                     "",
		Side:                     "",
		StopPrice:                "",
		IcebergQuantity:          "",
		Time:                     0,
		UpdateTime:               0,
		IsWorking:                false,
		IsIsolated:               false,
		OrigQuoteOrderQuantity:   "",
	}

	return order, nil
}

func (m *MockBinanceClient) GetOrderStatusError(symbol string, orderId int64) (*binance.Order, error) {
	return nil, errors.New("failed to get order status")
}

func (m *MockBinanceClient) ListOpenOrders(symbol string) ([]*binance.Order, error) {
	orders := []*binance.Order{
		&binance.Order{
			Symbol:                   symbol,
			OrderID:                  1234,
			OrderListId:              0,
			ClientOrderID:            "",
			Price:                    "",
			OrigQuantity:             "",
			ExecutedQuantity:         "",
			CummulativeQuoteQuantity: "",
			Status:                   "",
			TimeInForce:              "",
			Type:                     "",
			Side:                     "",
			StopPrice:                "",
			IcebergQuantity:          "",
			Time:                     0,
			UpdateTime:               0,
			IsWorking:                false,
			IsIsolated:               false,
			OrigQuoteOrderQuantity:   "",
		},
		&binance.Order{
			Symbol:                   symbol,
			OrderID:                  3456,
			OrderListId:              0,
			ClientOrderID:            "",
			Price:                    "",
			OrigQuantity:             "",
			ExecutedQuantity:         "",
			CummulativeQuoteQuantity: "",
			Status:                   "",
			TimeInForce:              "",
			Type:                     "",
			Side:                     "",
			StopPrice:                "",
			IcebergQuantity:          "",
			Time:                     0,
			UpdateTime:               0,
			IsWorking:                false,
			IsIsolated:               false,
			OrigQuoteOrderQuantity:   "",
		},
	}

	return orders, nil
}

func (m *MockBinanceClient) ListOpenOrdersError(symbol string) ([]*binance.Order, error) {
	return nil, errors.New("failed to get open orders")
}

func (m *MockBinanceClient) ListOrders(symbol string) ([]*binance.Order, error) {
	orders := []*binance.Order{
		{
			Symbol:                   symbol,
			OrderID:                  1234,
			OrderListId:              0,
			ClientOrderID:            "",
			Price:                    "",
			OrigQuantity:             "",
			ExecutedQuantity:         "",
			CummulativeQuoteQuantity: "",
			Status:                   "",
			TimeInForce:              "",
			Type:                     "",
			Side:                     "",
			StopPrice:                "",
			IcebergQuantity:          "",
			Time:                     0,
			UpdateTime:               0,
			IsWorking:                false,
			IsIsolated:               false,
			OrigQuoteOrderQuantity:   "",
		},
		{
			Symbol:                   symbol,
			OrderID:                  3456,
			OrderListId:              0,
			ClientOrderID:            "",
			Price:                    "",
			OrigQuantity:             "",
			ExecutedQuantity:         "",
			CummulativeQuoteQuantity: "",
			Status:                   "",
			TimeInForce:              "",
			Type:                     "",
			Side:                     "",
			StopPrice:                "",
			IcebergQuantity:          "",
			Time:                     0,
			UpdateTime:               0,
			IsWorking:                false,
			IsIsolated:               false,
			OrigQuoteOrderQuantity:   "",
		},
	}

	return orders, nil
}

func (m *MockBinanceClient) ListOrdersError(order string) ([]*binance.Order, error) {
	return nil, errors.New("failed to get orders")
}

func (m *MockBinanceClient) GetTradeHistory(symbol string) ([]*binance.TradeV3, error) {
	trades := []*binance.TradeV3{
		{
			ID:              1,
			Symbol:          symbol,
			OrderID:         1234,
			OrderListId:     1234,
			Price:           "0.200",
			Quantity:        "1",
			QuoteQuantity:   "1",
			Commission:      "1",
			CommissionAsset: "1",
			Time:            1234,
			IsBuyer:         true,
			IsMaker:         true,
			IsBestMatch:     true,
			IsIsolated:      true,
		},
		{
			ID:              2,
			Symbol:          symbol,
			OrderID:         1234,
			OrderListId:     1234,
			Price:           "0.200",
			Quantity:        "1",
			QuoteQuantity:   "1",
			Commission:      "1",
			CommissionAsset: "1",
			Time:            1234,
			IsBuyer:         true,
			IsMaker:         true,
			IsBestMatch:     true,
			IsIsolated:      true,
		},
	}

	return trades, nil
}

func (m *MockBinanceClient) GetTradeHistoryError(symbol string) ([]*binance.TradeV3, error) {
	return nil, errors.New("failed to get trade history")
}
