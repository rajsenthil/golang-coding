package algorithms

import "log"

func binary_search(input []int, key int, lo int, mid int, hi int) int {
	log.Printf("low: %d, mid: %d, high: %d", lo, mid, hi)
	if input[mid] == key {
		return mid
	}
	if lo > mid || hi <= mid {
		return -1
	}

	if key > input[mid] {
		lo = mid + 1
	} else {
		hi = mid
	}
	mid = (lo + hi) / 2
	return binary_search(input, key, lo, mid, hi)
}
