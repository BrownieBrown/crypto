package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockIndicators struct {
	mock.Mock
}

func (m *MockIndicators) RelativeStrengthIndex(close []float64, period int) []float64 {
	args := m.Called(close, period)
	return args.Get(0).([]float64)
}

func (m *MockIndicators) ExponentialMovingAverage(close []float64, period int) []float64 {
	args := m.Called(close, period)
	return args.Get(0).([]float64)
}

func (m *MockIndicators) BollingerBands(close []float64, period int, devUp float64, devDown float64) ([]float64, []float64, []float64) {
	args := m.Called(close, period, devUp, devDown)
	return args.Get(0).([]float64), args.Get(1).([]float64), args.Get(2).([]float64)
}

func (m *MockIndicators) FibonacciRetracement(high, low float64) map[string]float64 {
	args := m.Called(high, low)
	return args.Get(0).(map[string]float64)
}

func (m *MockIndicators) MovingAverageConvergenceDivergence(closePrices []float64, shortPeriod int, longPeriod int, signalPeriod int) ([]float64, []float64, []float64) {
	args := m.Called(closePrices, shortPeriod, longPeriod, signalPeriod)
	return args.Get(0).([]float64), args.Get(1).([]float64), args.Get(2).([]float64)
}

func (m *MockIndicators) AverageTrueRange(highPrices []float64, lowPrices []float64, closePrices []float64, timePeriod int) []float64 {
	args := m.Called(highPrices, lowPrices, closePrices, timePeriod)
	return args.Get(0).([]float64)
}

func (m *MockIndicators) VolumeWeightedAveragePrice(highPrices []float64, lowPrices []float64, closePrices []float64, volumes []float64) []float64 {
	args := m.Called(highPrices, lowPrices, closePrices, volumes)
	return args.Get(0).([]float64)
}
