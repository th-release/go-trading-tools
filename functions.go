package tradingtools

import (
	"math"
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
