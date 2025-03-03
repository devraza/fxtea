package fx

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

func Quadratic(a float64, b float64, c float64) []string {
	discriminant := math.Pow(b, 2) - 4*a*c

	var formattedRoots []string

	if discriminant < 0 {
		for _, v := range complexQuadratic(a, b, c) {
			sign := "+"
			if imag(v) < 0 {
				sign = "-"
			}

			formattedComplex := fmt.Sprintf(
				"%s %s %si",
				strconv.FormatFloat(Round(real(v), 0.001), 'f', -1, 64),
				sign,
				strconv.FormatFloat(Round(math.Abs(imag(v)), 0.001), 'f', -1, 64),
			)
			formattedRoots = append(formattedRoots, formattedComplex)
		}
		return formattedRoots
	} else {
		for _, v := range realQuadratic(a, b, c) {
			formattedRoots = append(formattedRoots, strconv.FormatFloat(v, 'f', -1, 64))
		}
		return formattedRoots
	}
}

func complexQuadratic(a float64, b float64, c float64) [2]complex128 {
	roots := [2]complex128{complex(0, 0), complex(0, 0)}

	discriminant := complex(b*b-4*a*c, 0)

	roots[0] = (-complex(b, 0) - cmplx.Sqrt(discriminant)) / complex(2*a, 0)
	roots[1] = (-complex(b, 0) + cmplx.Sqrt(discriminant)) / complex(2*a, 0)

	return roots
}

func realQuadratic(a float64, b float64, c float64) [2]float64 {
	roots := [2]float64{0., 0.}

	discriminant := b*b - 4*a*c

	roots[0] = Round(-b-(math.Sqrt(discriminant))/(2*a), 0.001)
	roots[1] = Round(-b+(math.Sqrt(discriminant))/(2*a), 0.001)

	return roots
}
