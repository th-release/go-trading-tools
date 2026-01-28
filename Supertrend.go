package tradingtools

type SupertrendResult struct {
	Supertrend []float64
	Direction  []int // 1 = uptrend, -1 = downtrend
}

// Supertrend calculates the Supertrend indicator
func (t *Tool) Supertrend(high, low, close []float64, length int, multiplier float64) SupertrendResult {
	n := len(close)
	atrValues := t.Atr(high, low, close, length)

	supertrend := makeSlice(n)
	direction := makeIntSlice(n)
	upperBand := makeSlice(n)
	lowerBand := makeSlice(n)

	for i := 0; i < n; i++ {
		hl2 := (high[i] + low[i]) / 2
		atrMult := multiplier * atrValues[i]
		upperBand[i] = hl2 + atrMult
		lowerBand[i] = hl2 - atrMult

		if i == 0 {
			supertrend[i] = upperBand[i]
			direction[i] = 1
			continue
		}

		// Adjust lower band
		if !(lowerBand[i] > lowerBand[i-1] || close[i-1] < lowerBand[i-1]) {
			lowerBand[i] = lowerBand[i-1]
		}

		// Adjust upper band
		if !(upperBand[i] < upperBand[i-1] || close[i-1] > upperBand[i-1]) {
			upperBand[i] = upperBand[i-1]
		}

		// Determine direction and supertrend value
		if supertrend[i-1] == upperBand[i-1] {
			if close[i] > upperBand[i] {
				direction[i] = 1
				supertrend[i] = lowerBand[i]
			} else {
				direction[i] = -1
				supertrend[i] = upperBand[i]
			}
		} else {
			if close[i] < lowerBand[i] {
				direction[i] = -1
				supertrend[i] = upperBand[i]
			} else {
				direction[i] = 1
				supertrend[i] = lowerBand[i]
			}
		}
	}

	return SupertrendResult{Supertrend: supertrend, Direction: direction}
}
