package tradingtools

// Atr calculates Average True Range
func (t *Tool) Atr(high, low, close []float64, length int) []float64 {
	return t.Rma(t.Tr(high, low, close), length)
}
