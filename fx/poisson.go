package fx

import (
  "math"
)

func factorial(n uint64) uint64 {
  if n == 0 {
      return 1
  }
  return n * factorial(n - 1)
}

func pow(a, b uint64) uint64 {
  n := a 
  for range b {
     n *= b
  }
  return n
}

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
