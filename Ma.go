package tradingtools

import "math"

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
