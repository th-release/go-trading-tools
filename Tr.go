package tradingtools

import "math"

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
