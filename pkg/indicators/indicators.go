package indicators

import "github.com/markcheno/go-talib"

type Indicators interface {
	RelativeStrengthIndex(close []float64, period int) []float64
	ExponentialMovingAverage(close []float64, period int) []float64
	BollingerBands(close []float64, period int, devUp float64, devDown float64) ([]float64, []float64, []float64)
	FibonacciRetracement(high, low float64) map[string]float64
	MovingAverageConvergenceDivergence(closePrices []float64, shortPeriod int, longPeriod int, signalPeriod int) ([]float64, []float64, []float64)
	AverageTrueRange(highPrices []float64, lowPrices []float64, closePrices []float64, timePeriod int) []float64
	VolumeWeightedAveragePrice(highPrices []float64, lowPrices []float64, closePrices []float64, volumes []float64) []float64
}

type TechnicalAnalysis struct{}

func (ta *TechnicalAnalysis) RelativeStrengthIndex(close []float64, period int) []float64 {
	return talib.Rsi(close, period)
}

func (ta *TechnicalAnalysis) ExponentialMovingAverage(close []float64, period int) []float64 {
	return talib.Ema(close, period)
}

func (ta *TechnicalAnalysis) BollingerBands(close []float64, period int, devUp float64, devDown float64) ([]float64, []float64, []float64) {
	upper, middle, lower := talib.BBands(close, period, devUp, devDown, 1)
	return upper, middle, lower
}

func (ta *TechnicalAnalysis) FibonacciRetracement(high, low float64) map[string]float64 {
	fibLevels := map[string]float64{
		"0.0":   high,
		"23.6":  high - 0.236*(high-low),
		"38.2":  high - 0.382*(high-low),
		"50.0":  high - 0.5*(high-low),
		"61.8":  high - 0.618*(high-low),
		"78.6":  high - 0.786*(high-low),
		"100.0": low,
	}

	return fibLevels
}

func (ta *TechnicalAnalysis) MovingAverageConvergenceDivergence(closePrices []float64, shortPeriod int, longPeriod int, signalPeriod int) ([]float64, []float64, []float64) {
	macd, signal, hist := talib.Macd(closePrices, shortPeriod, longPeriod, signalPeriod)
	return macd, signal, hist
}

func (ta *TechnicalAnalysis) AverageTrueRange(highPrices []float64, lowPrices []float64, closePrices []float64, timePeriod int) []float64 {
	atr := talib.Atr(highPrices, lowPrices, closePrices, timePeriod)
	return atr
}

func (ta *TechnicalAnalysis) VolumeWeightedAveragePrice(highPrices []float64, lowPrices []float64, closePrices []float64, volumes []float64) []float64 {
	typicalPrices := make([]float64, len(highPrices))
	for i := 0; i < len(highPrices); i++ {
		typicalPrices[i] = (highPrices[i] + lowPrices[i] + closePrices[i]) / 3
	}

	vwap := make([]float64, len(highPrices))
	cumulativeTPV := 0.0
	cumulativeVolume := 0.0
	for i := 0; i < len(highPrices); i++ {
		tpv := typicalPrices[i] * volumes[i]
		cumulativeTPV += tpv
		cumulativeVolume += volumes[i]
		vwap[i] = cumulativeTPV / cumulativeVolume
	}
	return vwap
}
