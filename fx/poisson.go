package fx

import (
	"math"
)

func PoissonPD(lambda float64, x uint64) float64 {
	return math.Pow(math.E, -lambda) * float64(math.Pow(lambda, float64(x))) / float64(factorial(x))
}

func PoissonCD(lambda float64, x uint64) float64 {
	cumulative := 0.
	var i uint64
	for i = 0; i <= x; i++ {
		cumulative += PoissonPD(lambda, i)
	}
	return cumulative
}
