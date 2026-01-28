package tradingtools

// Macd calculates the Moving Average Convergence Divergence
func (t *Tool) Macd(source []float64, fastLength, slowLength, signalLength int) MacdResult {
	n := len(source)
	fastEma := t.Ema(source, fastLength)
	slowEma := t.Ema(source, slowLength)

	macd := makeSlice(n)
	for i := 0; i < n; i++ {
		macd[i] = fastEma[i] - slowEma[i]
	}

	signal := t.Ema(macd, signalLength)

	histogram := makeSlice(n)
	for i := 0; i < n; i++ {
		histogram[i] = macd[i] - signal[i]
	}

	return MacdResult{Macd: macd, Signal: signal, Histogram: histogram}
}
