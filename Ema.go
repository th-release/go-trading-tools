package tradingtools

// Ema calculates Exponential Moving Average
// PineScript: ta.ema(source, length)
// alpha = 2 / (length + 1)
func (t *Tool) Ema(source []float64, length int) []float64 {
	alpha := 2.0 / float64(length+1)
	return t.expMovingAvg(source, alpha, length)
}
