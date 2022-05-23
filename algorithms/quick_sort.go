package algorithms

import (
	"log"
)

func quick_sort(input []int) []int {
	return quick_sort_helper(input, 0, len(input))
}

func quick_sort_helper(input []int, lo int, hi int) []int {
	log.Printf("Before quick_sort_helper input: %v, lo: %d, hi: %d", input, lo, hi)
	if lo < hi {
		pivot := quict_sort_partition(input, lo, len(input))
		log.Printf("After partitioning, the position of partition: %d", pivot)
		quick_sort_helper(input, lo, pivot)
		quick_sort_helper(input, pivot+1, hi)
	}
	log.Printf("After quick_sort_helper input: %v, lo: %d, hi: %d", input, lo, hi)
	return input
}

func quict_sort_partition(input []int, lo int, hi int) int {
	log.Printf("Before Partition: %v, with lo: %d and hi: %d", input, lo, hi)
	i := lo
	j := hi - 1
	pivot := input[lo]

	for i < j {
		for ; input[i] < pivot; i++ {
		}
		for ; input[j] > pivot; j-- {
		}
		//swap input[i] and input[j]
		if i < j {
			temp := input[j]
			input[j] = input[i]
			input[i] = temp
		}
	}
	log.Printf("After Partition: %v, with pivot: %d", input, j)
	return j
}
