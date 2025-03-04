package fx

import (
	"math/rand"
	"slices"
)

func QuickSort(list []int64) []int64 {
	if len(list) < 2 {
		return list
	}

	pivot := list[rand.Intn(len(list))]
	var lessThan, greaterThan []int64
	for _, v := range list {
		if v < pivot {
			lessThan = append(lessThan, v)
		} else if v > pivot {
			greaterThan = append(greaterThan, v)
		}
	}

	lessThan = QuickSort(lessThan)
	greaterThan = QuickSort(greaterThan)
	return slices.Concat(lessThan, []int64{pivot}, greaterThan)
}

func BinarySearch(list []int64, query int64) int64 {
	minimum := 0
	maximum := len(list) - 1

	for minimum <= maximum {
		pivot := minimum + (maximum-minimum)/2
		if list[pivot] == query {
			return int64(pivot)
		} else if list[pivot] < query {
			minimum = pivot + 1
		} else {
			maximum = pivot - 1
		}
	}

	return -1
}
