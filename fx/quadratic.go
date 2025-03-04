package fx

import (
	"fmt"
	"math"
	"math/cmplx"
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
				FormatFloat(Round(real(v), 0.0001)),
				sign,
				FormatFloat(Round(math.Abs(imag(v)), 0.0001)),
			)
			formattedRoots = append(formattedRoots, formattedComplex)
		}
		return formattedRoots
	} else {
		for _, v := range realQuadratic(a, b, c) {
			formattedRoots = append(formattedRoots, FormatFloat(v))
		}
		return formattedRoots
	}
}

func complexQuadratic(a float64, b float64, c float64) [2]complex128 {
	roots := [2]complex128{complex(0, 0), complex(0, 0)}

	discriminant := cmplx.Sqrt(complex(b*b-4*a*c, 0))

	roots[0] = (-complex(b, 0) - discriminant) / complex(2*a, 0)
	roots[1] = (-complex(b, 0) + discriminant) / complex(2*a, 0)

	return roots
}

func realQuadratic(a float64, b float64, c float64) [2]float64 {
	roots := [2]float64{0., 0.}

	discriminant := math.Sqrt(b*b - 4*a*c)

	roots[0] = Round((-b-(discriminant))/(2*a), 0.0001)
	roots[1] = Round((-b+(discriminant))/(2*a), 0.0001)

	return roots
}
