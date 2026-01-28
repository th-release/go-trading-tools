package tradingtools

// Ema calculates Exponential Moving Average
func (t *Tool) Ema(source []float64, length int) []float64 {
	alpha := 2.0 / float64(length+1)
	return t.expMovingAvg(source, alpha)
}
