package algorithms

import (
	"log"
	"testing"
)

func Test3WayPartition(t *testing.T) {
	input := []int{1, 2, 0, 1, 2, 0, 1, 2, 0}
	log.Printf("Given input: %v", input)
	result := three_way_partition(input, 1)
	log.Printf("After partitioning: %v", result)
}
