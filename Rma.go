package tradingtools

// Rma calculates Running Moving Average (Wilder's smoothing).
// PineScript: ta.rma(source, length)
// alpha = 1 / length
func (t *Tool) Rma(source []float64, length int) []float64 {
	alpha := 1.0 / float64(length)
	return t.expMovingAvg(source, alpha, length)
}
