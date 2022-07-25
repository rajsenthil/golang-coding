package dynamic_programming

import "log"

var lis_memoize []int
var lis_input_size int

func lis(input []int) int {
	n := len(input)
	lis_input_size = n
	lis_memoize = make([]int, n)
	for idx := 0; idx < len(lis_memoize); idx++ {
		lis_memoize[idx] = -1
	}
	lis_memoize[n-1] = 0
	log.Printf("Lis_size: %d", lis_input_size)
	lis_dyn(input)
	// the memoize does not include itself and so appending it at the last step
	// return lis_max_in_array(lis_memoize) + 1
	return lis_memoize[0] + 1
}

func lis_brute(input []int) [][]int {
	n := len(input)
	if n < 1 {
		return [][]int{{}}
	}
	log.Printf("Input param: %v", input)
	fi := input[0]
	oi := input[1:]
	result := lis_brute(oi)
	log.Printf("Result: %v, new char: %d", result, fi)
	combo := combine(result, fi)
	log.Printf("Result after combo: %v", result)
	log.Printf("Combo: %v", combo)
	return combo
}

func combine(in1 [][]int, new_char int) [][]int {
	if in1 == nil || len(in1) == 0 {
		return [][]int{{}, {new_char}}
	}
	buf := make([][]int, len(in1))
	for idx := range in1 {
		buf[idx] = make([]int, len(in1[idx]))
		copy(buf[idx], in1[idx])
		buf[idx] = append(buf[idx], new_char)
	}
	log.Printf("Input: %v, Output: %v", in1, buf)
	combos := append(in1, buf...)
	log.Printf("Input: %v, output: %v, Combos: %v ", in1, buf, combos)
	return combos
}

func lis_max_of(in1 int, in2 int) int {
	if in1 < in2 {
		return in2
	}
	return in1
}
func lis_dyn(input []int) {
	n := len(input)
	pos := lis_input_size - n
	log.Printf("Input: %v, size: %d, position: %d, memoize: %v", input, n, pos, lis_memoize)
	if n < 2 {
		return
	}

	if lis_memoize[pos] > -1 {
		return
	}
	for idx := 0; idx < len(input)-1; idx++ {
		lis_dyn(input[idx+1:])
		if input[idx] < input[idx+1] {
			lis_memoize[pos+idx] = 1 + lis_memoize[pos+idx+1]
		} else {
			next_best_pos := lis_find_next_greater_elem(input[idx], input)
			if next_best_pos > -1 {
				lis_memoize[pos+idx] = lis_memoize[pos+next_best_pos]
			} else {
				lis_memoize[pos+idx] = -1
			}
			// lis_memoize[pos+idx] = lis_memoize[pos+idx+1]
		}
	}
}

func lis_find_next_greater_elem(elem int, arr []int) int {
	next_best_pos := -1
	n := len(arr)
	for idx := 0; idx < n; idx++ {
		if elem < arr[idx] {
			next_best_pos = idx
			break
		}
	}
	log.Printf("Next best position: %d in array: %v, element: %d", next_best_pos, arr, elem)
	if next_best_pos == -1 {
		log.Printf("There is no next greater element found for: %d", elem)
		return -1
	}
	log.Printf("Next best position element: %d  with value: %d which is greater than: %d", arr[next_best_pos], next_best_pos, elem)
	return next_best_pos
}
