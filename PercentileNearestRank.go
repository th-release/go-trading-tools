package tradingtools

import (
	"math"
	"sort"
)

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
