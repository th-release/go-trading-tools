package tradingtools

// Rsi calculates Relative Strength Index
func (t *Tool) Rsi(source []float64, length int) []float64 {
	n := len(source)
	gains := makeSlice(n)
	losses := makeSlice(n)

	for i := 1; i < n; i++ {
		change := source[i] - source[i-1]
		if change > 0 {
			gains[i] = change
		} else {
			losses[i] = -change
		}
	}

	avgGain := t.Rma(gains, length)
	avgLoss := t.Rma(losses, length)

	result := makeSlice(n)
	for i := 0; i < n; i++ {
		if avgLoss[i] == 0 {
			result[i] = 100
		} else {
			rs := avgGain[i] / avgLoss[i]
			result[i] = 100 - (100 / (1 + rs))
		}
	}
	return result
}
