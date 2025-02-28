package fx

import (
	"math"
)

func PoissonPD(lambda float64, x uint64) float64 {
	return math.Pow(math.E, -lambda) * float64(pow(uint64(lambda), x)/factorial(x))
}

func PoissonCD(lambda float64, x uint64) float64 {
	cumulative := 0.
	for i := range x {
		cumulative += PoissonPD(lambda, i)
	}
	return cumulative
}
