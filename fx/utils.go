package fx

import (
	"fmt"
	"math"
	"strings"
)

func Round(n, nearest float64) float64 {
	return math.Round(n/nearest) * nearest
}

func FormatFloat(f float64) string {
	s := fmt.Sprintf("%.4f", f)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func pow(a, b uint64) uint64 {
	n := a
	for range b {
		n *= b
	}
	return n
}
