package fx

import (
	"fmt"
	"math"
)

func meanOfSlice(x []float64) float64 {
	var sum float64
	for _, i := range x {
		sum += i
	}
	return sum / float64(len(x))
}

var ErrTableMismatch = fmt.Errorf("columns of table are not the same length")

// Data should be linear, do not include outliers in the data
func PMCC(x []float64, y []float64) (float64, error) {
	var top float64
	var botParts [2]float64

	if len(x) != len(y) {
		return -2, ErrTableMismatch
	}

	meanX := meanOfSlice(x)
	meanY := meanOfSlice(y)

	for idx := range x {
		valX := x[idx] - meanX
		valY := y[idx] - meanY

		top += valX * valY
		botParts[0] += math.Pow(valX, 2)
		botParts[1] += math.Pow(valY, 2)
	}

	bot := math.Sqrt(botParts[0]) * math.Sqrt(botParts[1])
	return top / bot, nil
}
