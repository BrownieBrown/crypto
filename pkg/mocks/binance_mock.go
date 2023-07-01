package mocks

import (
	"crypto/pkg/constants"
	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/mock"
	"strconv"
)

type MockBinanceClient struct {
	mock.Mock
}

func (m *MockBinanceClient) GetAccountInfo() (*binance.Account, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*binance.Account), args.Error(1)
}

func (m *MockBinanceClient) GetCurrentPrice(symbol string) (float64, error) {
	args := m.Called(symbol)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockBinanceClient) PlaceOrder(symbol string, side string, quantity float64) (*binance.CreateOrderResponse, error) {
	args := m.Called(symbol, side, quantity)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*binance.CreateOrderResponse), args.Error(1)
}

func (m *MockBinanceClient) CancelOrder(symbol string, orderId int64) (*binance.CancelOrderResponse, error) {
	args := m.Called(symbol, orderId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*binance.CancelOrderResponse), args.Error(1)
}

func (m *MockBinanceClient) GetOrderStatus(symbol string, orderId int64) (*binance.Order, error) {
	args := m.Called(symbol, orderId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*binance.Order), args.Error(1)
}

func (m *MockBinanceClient) ListOpenOrders(symbol string) ([]*binance.Order, error) {
	args := m.Called(symbol)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*binance.Order), args.Error(1)
}

func (m *MockBinanceClient) ListOrders(symbol string) ([]*binance.Order, error) {
	args := m.Called(symbol)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*binance.Order), args.Error(1)
}

func (m *MockBinanceClient) GetTradeHistory(symbol string) ([]*binance.TradeV3, error) {
	args := m.Called(symbol)
	if args.Get(0) == nil {
		return nil, args.Error(1) // use 1 instead of 0
	}
	return args.Get(0).([]*binance.TradeV3), args.Error(1) // use 1 instead of 0
}

func (m *MockBinanceClient) GetHistoricalData(symbol string, interval string, limit int) ([]*binance.Kline, error) {
	args := m.Called(symbol, interval, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*binance.Kline), args.Error(1)
}

func NewMockData() *MockBinanceClient {
	mockClient := new(MockBinanceClient)
	mockClient.On("GetAccountInfo").Return(&binance.Account{
		MakerCommission:  constants.TestMakerComm,
		TakerCommission:  constants.TestMakerComm,
		BuyerCommission:  0,
		SellerCommission: 0,
		CanTrade:         true,
		CanWithdraw:      true,
		CanDeposit:       true,
		UpdateTime:       123456789,
		AccountType:      "SPOT",
		Balances: []binance.Balance{
			{
				Asset:  constants.TestAsset,
				Free:   constants.TestFree,
				Locked: "0.00000000",
			},
		},
		Permissions: []string{"SPOT"},
	}, nil)

	mockClient.On("GetCurrentPrice", constants.TestSymbol).Return(constants.TestPrice, nil)

	mockClient.On("PlaceOrder", constants.TestSymbol, constants.TestSide, constants.TestQuantity).Return(&binance.CreateOrderResponse{
		Symbol:                   constants.TestSymbol,
		OrderID:                  10,
		ClientOrderID:            "wOceeeOzNORyLiQfw7jd8S",
		TransactTime:             1507725176595,
		Price:                    "0.02000000",
		OrigQuantity:             strconv.FormatFloat(constants.TestQuantity, 'f', 6, 64),
		ExecutedQuantity:         "0.000000002",
		CummulativeQuoteQuantity: "0.000000002",
		IsIsolated:               false,
		Status:                   "NEW",
		TimeInForce:              "GTC",
		Type:                     "LIMIT",
		Side:                     constants.TestSide,
	}, nil)

	mockClient.On("CancelOrder", constants.TestSymbol, constants.TestOrderID).Return(&binance.CancelOrderResponse{
		Symbol:                   constants.TestSymbol,
		OrigClientOrderID:        "1",
		OrderID:                  constants.TestOrderID,
		OrderListID:              0,
		ClientOrderID:            "1234",
		TransactTime:             0,
		Price:                    "0.02000000",
		OrigQuantity:             "1",
		ExecutedQuantity:         "1",
		CummulativeQuoteQuantity: "1",
		Status:                   "NEW",
		TimeInForce:              "1234",
		Type:                     "test",
		Side:                     constants.TestSide,
	}, nil)

	mockClient.On("GetOrderStatus", constants.TestSymbol, constants.TestOrderID).Return(&binance.Order{
		Symbol:                   constants.TestSymbol,
		OrderID:                  constants.TestOrderID,
		OrderListId:              0,
		ClientOrderID:            "myOrder1",
		Price:                    "0.1",
		OrigQuantity:             "1.0",
		ExecutedQuantity:         "0.000000002",
		CummulativeQuoteQuantity: "0.000000002",
		Status:                   "NEW",
		TimeInForce:              "GTC",
		Type:                     "LIMIT",
		Side:                     "BUY",
		StopPrice:                "0.0",
		IcebergQuantity:          "0.0",
		Time:                     1499827319559,
		UpdateTime:               1499827319559,
		IsWorking:                true,
		OrigQuoteOrderQuantity:   "0.000000",
	}, nil)

	mockClient.On("ListOpenOrders", constants.TestSymbol).Return([]*binance.Order{
		{
			Symbol:                   constants.TestSymbol,
			OrderID:                  1,
			OrderListId:              0,
			ClientOrderID:            "myOrder1",
			Price:                    "0.1",
			OrigQuantity:             "1.0",
			ExecutedQuantity:         "0.000000002",
			CummulativeQuoteQuantity: "0.000000002",
			Status:                   "NEW",
			TimeInForce:              "GTC",
			Type:                     "LIMIT",
			Side:                     "BUY",
			StopPrice:                "0.0",
			IcebergQuantity:          "0.0",
			Time:                     1499827319559,
			UpdateTime:               1499827319559,
			IsWorking:                true,
			OrigQuoteOrderQuantity:   "0.000000",
		},
		{
			Symbol:                   constants.TestSymbol,
			OrderID:                  2,
			OrderListId:              0,
			ClientOrderID:            "myOrder1",
			Price:                    "0.1",
			OrigQuantity:             "1.0",
			ExecutedQuantity:         "0.000000002",
			CummulativeQuoteQuantity: "0.000000002",
			Status:                   "NEW",
			TimeInForce:              "GTC",
			Type:                     "LIMIT",
			Side:                     "BUY",
			StopPrice:                "0.0",
			IcebergQuantity:          "0.0",
			Time:                     1499827319559,
			UpdateTime:               1499827319559,
			IsWorking:                true,
			OrigQuoteOrderQuantity:   "0.000000",
		},
	}, nil)

	mockClient.On("ListOrders", constants.TestSymbol).Return([]*binance.Order{
		{
			Symbol:                   constants.TestSymbol,
			OrderID:                  1,
			OrderListId:              0,
			ClientOrderID:            "myOrder1",
			Price:                    "0.1",
			OrigQuantity:             "1.0",
			ExecutedQuantity:         "0.000000002",
			CummulativeQuoteQuantity: "0.000000002",
			Status:                   "NEW",
			TimeInForce:              "GTC",
			Type:                     "LIMIT",
			Side:                     "BUY",
			StopPrice:                "0.0",
			IcebergQuantity:          "0.0",
			Time:                     1499827319559,
			UpdateTime:               1499827319559,
			IsWorking:                true,
			OrigQuoteOrderQuantity:   "0.000000",
		},
		{
			Symbol:                   constants.TestSymbol,
			OrderID:                  2,
			OrderListId:              0,
			ClientOrderID:            "myOrder1",
			Price:                    "0.1",
			OrigQuantity:             "1.0",
			ExecutedQuantity:         "0.000000002",
			CummulativeQuoteQuantity: "0.000000002",
			Status:                   "NEW",
			TimeInForce:              "GTC",
			Type:                     "LIMIT",
			Side:                     "BUY",
			StopPrice:                "0.0",
			IcebergQuantity:          "0.0",
			Time:                     1499827319559,
			UpdateTime:               1499827319559,
			IsWorking:                true,
			OrigQuoteOrderQuantity:   "0.000000",
		},
	}, nil)

	mockClient.On("GetTradeHistory", constants.TestSymbol).Return([]*binance.TradeV3{
		{
			ID:              1,
			Symbol:          constants.TestSymbol,
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
			ID:              1,
			Symbol:          constants.TestSymbol,
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
	}, nil)

	mockClient.On("GetHistoricalData", constants.TestSymbol, constants.TestInterval, constants.TestLimit).Return([]*binance.Kline{
		{
			OpenTime:                 1499040000000,
			Open:                     "0.01634790",
			High:                     "0.80000000",
			Low:                      "0.01575800",
			Close:                    "0.01577100",
			Volume:                   "148976.11427815",
			CloseTime:                1499644799999,
			QuoteAssetVolume:         "2434.19055334",
			TradeNum:                 308,
			TakerBuyBaseAssetVolume:  "1756.87402397",
			TakerBuyQuoteAssetVolume: "28.46694368",
		},
		{
			OpenTime:                 1499040000000,
			Open:                     "0.01634790",
			High:                     "0.80000000",
			Low:                      "0.01575800",
			Close:                    "0.01577100",
			Volume:                   "148976.11427815",
			CloseTime:                1499644799999,
			QuoteAssetVolume:         "2434.19055334",
			TradeNum:                 308,
			TakerBuyBaseAssetVolume:  "1756.87402397",
			TakerBuyQuoteAssetVolume: "28.46694368",
		},
	}, nil)

	return mockClient
}
