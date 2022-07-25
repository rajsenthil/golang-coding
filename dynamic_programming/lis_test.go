package dynamic_programming

import (
	"log"
	"testing"
)

func TestLIS(t *testing.T) {
	input := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	// input = []int{3, 10, 2, 1, 20}
	// input = []int{3, 2}
	input = []int{50, 3, 10, 7, 40, 80}
	input = []int{4, 10, 4, 3, 8, 9}
	result := lis(input)
	log.Printf("The longest incresing subsequence of input: %v is: %d", input, result)
}

func TestLISBrute(t *testing.T) {
	input := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	input = []int{3, 10, 2, 1, 20}
	// input = []int{3, 2}
	// input = []int{50, 3, 10, 7, 40, 80}
	input = []int{4, 10, 4, 3, 8, 9}
	// input = []int{0, 1, 0, 3, 2, 3}
	input = []int{1, 2, 3}
	result := lis_brute(input)
	log.Printf("The longest incresing subsequence of input: %v is: %v", input, result)
}

func TestNextGreaterElem(t *testing.T) {
	input := []int{4, 10, 4, 3, 8, 9}
	input = []int{4, 3, 8, 9}
	elem := 4
	result := lis_find_next_greater_elem(elem, input)
	log.Printf("Next greater element for %d in the array: %d is %d", elem, input, result)
}

func TestCombine(t *testing.T) {
	input := [][]int{{1, 2}, {3, 4, 5}, {6, 7, 8, 9}}
	new_char := 10
	log.Printf("Input: %v, new char: %d", input, new_char)
	result := combine(input, new_char)
	log.Printf("Combination: %v", result)
}

func TestCombineNil(t *testing.T) {
	input := [][]int{}
	new_char := 10
	log.Printf("Input: %v, new char: %d", input, new_char)
	result := combine(input, new_char)
	log.Printf("Combination: %v", result)
}

/**
[10 22 9 33 21 50 41 60 80]
[5   4 4  3  3  2  2  1  0]

[4 10 4 3 8 9]
[3  2 2 2 1 0]

**/
