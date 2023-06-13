package api

import (
	"context"
	"crypto/pkg/logger"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"go.uber.org/zap"
	"strconv"
)

type BinanceAPI interface {
	GetAccountInfo() (*binance.Account, error)
	GetCurrentPrice(symbol string) (float64, error)
	PlaceOrder(symbol string, side string, quantity float64) (*binance.CreateOrderResponse, error)
	CancelOrder(symbol string, orderId int64) (*binance.CancelOrderResponse, error)
	GetOrderStatus(symbol string, orderId int64) (*binance.Order, error)
	ListOpenOrders(symbol string) ([]*binance.Order, error)
	ListOrders(symbol string) ([]*binance.Order, error)
	GetTradeHistory(symbol string) ([]*binance.TradeV3, error)
}

type BinanceClient struct {
	Client *binance.Client
}

var _ BinanceAPI = &BinanceClient{}

func NewBinanceClient(apiKey string, secretKey string) *BinanceClient {
	return &BinanceClient{
		Client: binance.NewClient(apiKey, secretKey),
	}
}

func (bc *BinanceClient) GetAccountInfo() (*binance.Account, error) {
	account, err := bc.Client.NewGetAccountService().Do(context.Background())
	if err != nil {
		logger.Log.Error("Failed to get account info", zap.Error(err))
		return nil, err
	}

	return account, nil
}

func (bc *BinanceClient) GetCurrentPrice(symbol string) (float64, error) {
	prices, err := bc.Client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		logger.Log.Error("Failed to fetch prices for symbol", zap.String("symbol", symbol), zap.Error(err))
		return 0, err
	}

	if len(prices) == 0 {
		logger.Log.Error("No prices found for symbol", zap.String("symbol", symbol), zap.Error(err))
		return 0, err
	}

	price, err := strconv.ParseFloat(prices[0].Price, 64)
	if err != nil {
		logger.Log.Error("Failed to parse price for crypto currency", zap.Error(err))
		return 0, err
	}

	return price, nil
}

func (bc *BinanceClient) PlaceOrder(symbol string, side string, quantity float64) (*binance.CreateOrderResponse, error) {
	order, err := bc.Client.NewCreateOrderService().
		Symbol(symbol).
		Side(binance.SideType(side)).
		Type(binance.OrderTypeMarket).
		Quantity(fmt.Sprintf("%.8f", quantity)).
		Do(context.Background())

	if err != nil {
		logger.Log.Error("Error placing order", zap.String("symnbol", symbol), zap.String("side", side), zap.Int("quantity", int(quantity)), zap.Error(err))
		return nil, err
	}

	return order, nil
}

func (bc *BinanceClient) CancelOrder(symbol string, orderId int64) (*binance.CancelOrderResponse, error) {
	response, err := bc.Client.NewCancelOrderService().Symbol(symbol).OrderID(orderId).Do(context.Background())

	if err != nil {
		logger.Log.Error("Failed to cancel order", zap.String("symbol", symbol), zap.Int("orderId", int(orderId)), zap.Error(err))
		return nil, err
	}

	return response, nil
}

func (bc *BinanceClient) GetOrderStatus(symbol string, orderId int64) (*binance.Order, error) {
	order, err := bc.Client.NewGetOrderService().
		Symbol(symbol).
		OrderID(orderId).
		Do(context.Background())

	if err != nil {
		logger.Log.Error("Failed to get order status", zap.String("symbol", symbol), zap.Int("orderId", int(orderId)), zap.Error(err))
		return nil, err
	}

	return order, nil
}

func (bc *BinanceClient) ListOpenOrders(symbol string) ([]*binance.Order, error) {
	orders, err := bc.Client.NewListOpenOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		logger.Log.Error("Failed to get open orders", zap.String("symbol", symbol), zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (bc *BinanceClient) ListOrders(symbol string) ([]*binance.Order, error) {
	orders, err := bc.Client.NewListOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		logger.Log.Error("Failed to get orders", zap.String("symbol", symbol), zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (bc *BinanceClient) GetTradeHistory(symbol string) ([]*binance.TradeV3, error) {
	trades, err := bc.Client.NewListTradesService().
		Symbol(symbol).
		Do(context.Background())

	if err != nil {
		logger.Log.Error("Failed to get trade history", zap.String("symbol", symbol), zap.Error(err))
		return nil, err
	}

	return trades, nil
}
