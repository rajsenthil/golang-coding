package algorithms

import "log"

var count = 0
var input []int

func permutation(param []int) {
	// size := len(param)
	input = param
	permutation_helper(param)
	log.Printf("Count: %d", count)
}

func permutation_helper(param []int) {
	log.Printf("Param: %v", param)
	size := len(param)
	if size == 1 {
		return
	}

	for idx := 0; idx < size; idx++ {
		permutation_helper(param[1:])
		if size == 2 {
			log.Printf("%v", param)
			count++
		}
		rotate(param)
	}
}

func rotate(param []int) []int {
	size := len(param)
	// log.Printf("Rotate: %v", param)
	if size <= 1 {
		return param
	}
	first := param[0]
	for idx := 0; idx < size-1; idx++ {
		param[idx] = param[idx+1]
	}
	param[size-1] = first
	// log.Printf("RotateD: %v", param)
	return param
}
