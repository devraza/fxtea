package fx

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
