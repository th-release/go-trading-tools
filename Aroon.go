package tradingtools

type AroonResult struct {
	Up   []float64
	Down []float64
}

// Aroon calculates the Aroon indicator
// PineScript: ta.aroon(high, low, length)
// Aroon Up = 100 * (length - bars since highest high) / length
// Aroon Down = 100 * (length - bars since lowest low) / length
func (t *Tool) Aroon(high, low []float64, length int) AroonResult {
	n := len(high)
	up := makeSlice(n)
	down := makeSlice(n)
	lengthF := float64(length)

	for i := 0; i < n; i++ {
		if i < length {
			continue
		}

		barsSinceHigh, barsSinceLow := 0, 0
		highestVal, lowestVal := high[i], low[i]

		for j := 0; j <= length; j++ {
			idx := i - j
			if high[idx] >= highestVal {
				highestVal = high[idx]
				barsSinceHigh = j
			}
			if low[idx] <= lowestVal {
				lowestVal = low[idx]
				barsSinceLow = j
			}
		}

		up[i] = 100 * (lengthF - float64(barsSinceHigh)) / lengthF
		down[i] = 100 * (lengthF - float64(barsSinceLow)) / lengthF
	}

	return AroonResult{Up: up, Down: down}
}
