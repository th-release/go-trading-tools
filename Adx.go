package tradingtools

import "math"

type AdxResult struct {
	Adx     []float64
	PlusDi  []float64
	MinusDi []float64
}

// Adx calculates the Average Directional Index
func (t *Tool) Adx(high, low, close []float64, length int) AdxResult {
	n := len(close)
	trValues := t.Tr(high, low, close)

	plusDm := makeSlice(n)
	minusDm := makeSlice(n)

	for i := 1; i < n; i++ {
		upMove := high[i] - high[i-1]
		downMove := low[i-1] - low[i]

		if upMove > downMove && upMove > 0 {
			plusDm[i] = upMove
		}
		if downMove > upMove && downMove > 0 {
			minusDm[i] = downMove
		}
	}

	smoothedTr := t.Rma(trValues, length)
	smoothedPlusDm := t.Rma(plusDm, length)
	smoothedMinusDm := t.Rma(minusDm, length)

	plusDi := makeSlice(n)
	minusDi := makeSlice(n)
	dx := makeSlice(n)

	for i := 0; i < n; i++ {
		plusDi[i] = safeDiv(100*smoothedPlusDm[i], smoothedTr[i], 0)
		minusDi[i] = safeDiv(100*smoothedMinusDm[i], smoothedTr[i], 0)
		diSum := plusDi[i] + minusDi[i]
		dx[i] = safeDiv(100*math.Abs(plusDi[i]-minusDi[i]), diSum, 0)
	}

	return AdxResult{Adx: t.Rma(dx, length), PlusDi: plusDi, MinusDi: minusDi}
}
