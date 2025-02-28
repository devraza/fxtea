package fx

import "math"

func Sum(o float64, e float64) float64 {
  return math.Pow((o-e), 2)/e
}
