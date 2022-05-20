package algorithms

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// input := []int{12, 11, 13, 5, 7, 6}
	input := []int{3, 4, 2, 1, 7, 5, 8, 9, 0, 6}
	log.Printf("Input va;ues: %v", input)
	merge := Merge{
		input: input,
	}
	result := merge.merge_sort()
	log.Printf("Sorted result: %v", result)
}
