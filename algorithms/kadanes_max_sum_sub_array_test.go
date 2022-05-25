package algorithms

import (
	"log"
	"testing"
)

func TestKadanes(t *testing.T) {
	input := []int{-2, 3, 2, -1}
	result := kadanes(input)
	log.Printf("For the given array: %v, the max sub array: %d", input, result)
}
