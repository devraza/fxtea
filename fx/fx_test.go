package fx

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	unsorted := []int64{45, 28, 97, 100, 99, 78}
	sorted := []int64{28, 45, 78, 97, 99, 100}

	res := QuickSort(unsorted)

	if !slices.Equal(res, sorted) {
		t.Errorf(`QuickSort([]int64{45, 28, 97, 100, 99, 78}) = %v, want match for %v`, res, sorted)
	}
}
