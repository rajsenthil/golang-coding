package algorithms

import "log"

func first_missing_positive_with_hash(nums []int) int {
	min := 1
	num_pair := make(map[int]bool)
	for idx := 0; idx < len(nums); idx++ {
		num_pair[nums[idx]] = true
	}

	log.Printf("Num pair: %v", num_pair)

	for min = 1; min < len(nums)+1; min++ {
		if _, ok := num_pair[min]; ok {
			log.Printf("num: %d exists", min)
			continue
		} else {
			log.Printf("Minimum found in the loop: %d", min)
			return min
		}
	}
	return min
}
