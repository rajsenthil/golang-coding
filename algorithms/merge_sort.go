package algorithms

import "log"

type Merge struct {
	input []int
}

func (m *Merge) merge_sort() []int {
	return merge_sort_helper(m.input)
}

func merge_sort_helper(input []int) []int {
	log.Printf("Input: %v", input)
	// Step 1: split
	if len(input) < 2 {
		return (input)
	}
	mid := len(input) / 2
	lValues := merge_sort_helper(input[:mid])
	rValues := merge_sort_helper(input[mid:])
	log.Printf("Lvalues: %v, RValues: %v", lValues, rValues)
	// Step 2: Merge
	return merge_1(lValues, rValues)
}

func merge_1(left, right []int) []int {
	log.Printf("Left array: %v, right array: %v", left, right)
	buf := make([]int, len(left)+len(right))
	i := 0
	j := 0
	k := 0

	for ; i < len(left) && j < len(right); k++ {
		if left[i] < right[j] {
			buf[k] = left[i]
			i++
		} else {
			buf[k] = right[j]
			j++
		}
	}

	for ; i < len(left); i++ {
		buf[k] = left[i]
		k++
	}
	for ; j < len(right); j++ {
		buf[k] = right[j]
		k++
	}
	return buf
}
