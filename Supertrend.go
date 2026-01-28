package tradingtools

import "math"

type SupertrendResult struct {
	Supertrend []float64
	Direction  []int // 1 = uptrend, -1 = downtrend
}

// Supertrend calculates the Supertrend indicator
// PineScript: ta.supertrend(multiplier, length)
func (t *Tool) Supertrend(high, low, close []float64, length int, multiplier float64) SupertrendResult {
	n := len(close)
	atrValues := t.Atr(high, low, close, length)

	supertrend := makeSlice(n)
	direction := makeIntSlice(n)
	upperBand := makeSlice(n)
	lowerBand := makeSlice(n)

	// Find first valid index (where ATR is not NaN)
	firstValid := -1
	for i := 0; i < n; i++ {
		if !math.IsNaN(atrValues[i]) {
			firstValid = i
			break
		}
	}

	for i := 0; i < n; i++ {
		if math.IsNaN(atrValues[i]) {
			supertrend[i] = math.NaN()
			upperBand[i] = math.NaN()
			lowerBand[i] = math.NaN()
			direction[i] = 0
			continue
		}

		hl2 := (high[i] + low[i]) / 2
		atrMult := multiplier * atrValues[i]
		upperBand[i] = hl2 + atrMult
		lowerBand[i] = hl2 - atrMult

		if i == firstValid {
			supertrend[i] = upperBand[i]
			direction[i] = 1
			continue
		}

		prevUpper := upperBand[i-1]
		prevLower := lowerBand[i-1]

		// Adjust lower band
		if !math.IsNaN(prevLower) && !(lowerBand[i] > prevLower || close[i-1] < prevLower) {
			lowerBand[i] = prevLower
		}

		// Adjust upper band
		if !math.IsNaN(prevUpper) && !(upperBand[i] < prevUpper || close[i-1] > prevUpper) {
			upperBand[i] = prevUpper
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
