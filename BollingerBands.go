package tradingtools

import "math"

type BollingerBandsResult struct {
	Upper  []float64
	Middle []float64
	Lower  []float64
}

// BollingerBands calculates the Bollinger Bands indicator
func (t *Tool) BollingerBands(source []float64, length int, mult float64) BollingerBandsResult {
	n := len(source)
	middle := t.Ma(source, length)
	upper := makeSlice(n)
	lower := makeSlice(n)

	for i := 0; i < n; i++ {
		if i < length-1 {
			upper[i] = math.NaN()
			lower[i] = math.NaN()
			continue
		}

		stdev := math.Sqrt(varianceWindow(source, i, length, middle[i]))
		band := mult * stdev
		upper[i] = middle[i] + band
		lower[i] = middle[i] - band
	}

	return BollingerBandsResult{Upper: upper, Middle: middle, Lower: lower}
}
