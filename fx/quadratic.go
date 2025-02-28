package fx

import (
  "math"
)

func Quadratic(a float64, b float64, c float64) [2]float64 {
  roots := [2]float64{0., 0.}

  roots[0] = Round(-b - (math.Sqrt((math.Pow(b, 2) - 4*a*c)))/(2*a), 0.001)
  roots[1] = Round(-b + (math.Sqrt((math.Pow(b, 2) - 4*a*c)))/(2*a), 0.001)

  return roots
}
