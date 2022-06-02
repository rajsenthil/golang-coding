package algorithms

import (
	"log"
	"testing"
)

func TestSortedMedianArrays(t *testing.T) {
	array1 := []int{1, 3}
	array2 := []int{2}
	array1 = []int{-5, 3, 6, 12, 15}
	array2 = []int{-12, -10, -6, -3, 4, 10}
	array1 = []int{2, 3, 5, 8}
	array2 = []int{10, 12, 14, 16, 18, 20}
	array1 = []int{1, 3}
	array2 = []int{2}
	median, part_x := sorted_arrays_median(array1, array2)
	log.Printf("The median of array1: %v and array2: %v is %d and the partition_x: %d", array1, array2, median, part_x)
}
