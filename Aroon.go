package tradingtools

type AroonResult struct {
	Up   []float64
	Down []float64
}

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
