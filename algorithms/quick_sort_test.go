package algorithms

import (
	"log"
	"testing"
)

func TestQuickSortPartition(t *testing.T) {
	// input := []int{12, 9, 6, 13, 8, 19, 23, 4}
	input := []int{9, 6, 8, 4, 12, 19, 23, 13}
	// input = []int{4, 9, 6, 8, 12, 19, 23, 13, 2147483647}
	input = []int{8, 7, 6, 1, 0, 9, 2}
	result := quict_sort_partition(input, 0, len(input))
	log.Printf("The position of the pivot: %d", result)
}

func TestQuickSort(t *testing.T) {
	input := []int{12, 9, 6, 13, 8, 19, 23, 4}
	input = []int{8, 7, 6, 1, 0, 9, 2}
	result := quick_sort(input)
	log.Printf("Sorted Result: %v", result)
}
