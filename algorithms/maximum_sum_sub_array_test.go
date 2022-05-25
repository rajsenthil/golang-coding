package algorithms

import (
	"log"
	"testing"
)

func TestMaxSumSubArray(t *testing.T) {
	//For example, for the array of values [-2, 1, -3, 4, -1, 2, 1, -5, 4], the contiguous subarray with the largest sum is [4, −1, 2, 1], with sum 6.
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	input = []int{-2, -3, 4, -1, -2, 1, 5, -3}
	result, startPos, endPos := maxSumSubArray(input, 0, len(input))
	log.Printf("Result: %d", result)
	log.Printf("For the given input: %v, the start index: %d, end index: %d, the result is %d", input, startPos, endPos, result)
}
