package fx

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

func ChiSum(o float64, e float64) float64 {
	return math.Pow((o-e), 2) / e
}

func ChiCritical(df float64, alpha float64) float64 {
	chi := &distuv.ChiSquared{
		K:   df,
		Src: nil,
	}
	return chi.Quantile(1 - alpha)
}
