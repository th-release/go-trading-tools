package tradingtools

import (
	"math"
	"sort"
)

// ============================================================================
// Helper functions (internal)
// ============================================================================

// makeSlice creates a new float64 slice of given size
func makeSlice(n int) []float64 {
	return make([]float64, n)
}

// makeIntSlice creates a new int slice of given size
func makeIntSlice(n int) []int {
	return make([]int, n)
}

// expMovingAvg calculates exponential moving average with given alpha
func (t *Tool) expMovingAvg(source []float64, alpha float64) []float64 {
	n := len(source)
	result := makeSlice(n)

	for i := 0; i < n; i++ {
		if i == 0 {
			result[i] = source[i]
		} else {
			result[i] = alpha*source[i] + (1-alpha)*result[i-1]
		}
	}
	return result
}

// sumWindow calculates sum of source[i-length+1:i+1]
func sumWindow(source []float64, i, length int) float64 {
	sum := 0.0
	for j := 0; j < length; j++ {
		sum += source[i-j]
	}
	return sum
}

// varianceWindow calculates variance of source[i-length+1:i+1] given mean
func varianceWindow(source []float64, i, length int, mean float64) float64 {
	variance := 0.0
	for j := 0; j < length; j++ {
		diff := source[i-j] - mean
		variance += diff * diff
	}
	return variance / float64(length)
}

// copyWindow copies source[i-length+1:i+1] to a new slice
func copyWindow(source []float64, i, length int) []float64 {
	window := makeSlice(length)
	copy(window, source[i-length+1:i+1])
	return window
}

// max3 returns the maximum of three float64 values
func max3(a, b, c float64) float64 {
	return math.Max(a, math.Max(b, c))
}

// safeDiv returns a/b, or fallback if b is zero
func safeDiv(a, b, fallback float64) float64 {
	if b == 0 {
		return fallback
	}
	return a / b
}

// ============================================================================
// Result types
// ============================================================================

type AroonResult struct {
	Up   []float64
	Down []float64
}

type MacdResult struct {
	Macd      []float64
	Signal    []float64
	Histogram []float64
}

// ============================================================================
// Moving Averages
// ============================================================================

// Ma calculates Simple Moving Average
func (t *Tool) Ma(source []float64, length int) []float64 {
	n := len(source)
	result := makeSlice(n)
	lengthF := float64(length)

	for i := 0; i < n; i++ {
		if i < length-1 {
			result[i] = math.NaN()
			continue
		}
		result[i] = sumWindow(source, i, length) / lengthF
	}
	return result
}

// Ema calculates Exponential Moving Average
func (t *Tool) Ema(source []float64, length int) []float64 {
	alpha := 2.0 / float64(length+1)
	return t.expMovingAvg(source, alpha)
}

// Rma calculates Running Moving Average (Wilder's smoothing)
func (t *Tool) Rma(source []float64, length int) []float64 {
	alpha := 1.0 / float64(length)
	return t.expMovingAvg(source, alpha)
}

// ============================================================================
// Volatility Indicators
// ============================================================================

// Tr calculates True Range
func (t *Tool) Tr(high, low, close []float64) []float64 {
	n := len(high)
	result := makeSlice(n)

	for i := 0; i < n; i++ {
		hl := high[i] - low[i]
		if i == 0 {
			result[i] = hl
		} else {
			hc := math.Abs(high[i] - close[i-1])
			lc := math.Abs(low[i] - close[i-1])
			result[i] = max3(hl, hc, lc)
		}
	}
	return result
}

// Atr calculates Average True Range
func (t *Tool) Atr(high, low, close []float64, length int) []float64 {
	return t.Rma(t.Tr(high, low, close), length)
}

// Stdev calculates the standard deviation over a rolling window
func (t *Tool) Stdev(source []float64, length int) []float64 {
	n := len(source)
	result := makeSlice(n)
	ma := t.Ma(source, length)

	for i := 0; i < n; i++ {
		if i < length-1 {
			result[i] = 0
			continue
		}
		result[i] = math.Sqrt(varianceWindow(source, i, length, ma[i]))
	}
	return result
}

// ============================================================================
// Momentum Indicators
// ============================================================================

// Rsi calculates Relative Strength Index
func (t *Tool) Rsi(source []float64, length int) []float64 {
	n := len(source)
	gains := makeSlice(n)
	losses := makeSlice(n)

	for i := 1; i < n; i++ {
		change := source[i] - source[i-1]
		if change > 0 {
			gains[i] = change
		} else {
			losses[i] = -change
		}
	}

	avgGain := t.Rma(gains, length)
	avgLoss := t.Rma(losses, length)

	result := makeSlice(n)
	for i := 0; i < n; i++ {
		if avgLoss[i] == 0 {
			result[i] = 100
		} else {
			rs := avgGain[i] / avgLoss[i]
			result[i] = 100 - (100 / (1 + rs))
		}
	}
	return result
}

// Macd calculates the Moving Average Convergence Divergence
func (t *Tool) Macd(source []float64, fastLength, slowLength, signalLength int) MacdResult {
	n := len(source)
	fastEma := t.Ema(source, fastLength)
	slowEma := t.Ema(source, slowLength)

	macd := makeSlice(n)
	for i := 0; i < n; i++ {
		macd[i] = fastEma[i] - slowEma[i]
	}

	signal := t.Ema(macd, signalLength)

	histogram := makeSlice(n)
	for i := 0; i < n; i++ {
		histogram[i] = macd[i] - signal[i]
	}

	return MacdResult{Macd: macd, Signal: signal, Histogram: histogram}
}

// ============================================================================
// Trend Indicators
// ============================================================================

// Aroon calculates the Aroon indicator
func (t *Tool) Aroon(high, low []float64, length int) AroonResult {
	n := len(high)
	up := makeSlice(n)
	down := makeSlice(n)
	lengthF := float64(length)

	for i := 0; i < n; i++ {
		if i < length {
			continue
		}

		highestIdx, lowestIdx := 0, 0
		highestVal, lowestVal := high[i-length], low[i-length]

		for j := 0; j <= length; j++ {
			idx := i - length + j
			if high[idx] >= highestVal {
				highestVal = high[idx]
				highestIdx = j
			}
			if low[idx] <= lowestVal {
				lowestVal = low[idx]
				lowestIdx = j
			}
		}

		up[i] = 100 * float64(highestIdx) / lengthF
		down[i] = 100 * float64(lowestIdx) / lengthF
	}

	return AroonResult{Up: up, Down: down}
}

// ============================================================================
// Statistical Functions
// ============================================================================

// PercentileNearestRank calculates the percentile using the nearest rank method
func (t *Tool) PercentileNearestRank(source []float64, length int, percentile int) []float64 {
	n := len(source)
	result := makeSlice(n)
	percentileF := float64(percentile) / 100.0

	for i := 0; i < n; i++ {
		if i < length-1 {
			continue
		}

		window := copyWindow(source, i, length)
		sort.Float64s(window)

		index := int(math.Ceil(percentileF*float64(length))) - 1
		if index < 0 {
			index = 0
		}
		result[i] = window[index]
	}

	return result
}
