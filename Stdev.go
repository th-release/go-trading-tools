package tradingtools

import "math"

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
