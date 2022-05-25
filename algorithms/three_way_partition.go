package algorithms

import "log"

func three_way_partition(input []int, mid int) []int {
	i := 0
	j := 0
	k := len(input) - 1

	for j <= k {
		log.Printf("i: %d, j: %d, k: %d, input[j]: %d", i, j, k, input[j])
		if input[j] < mid {
			input[i], input[j] = swap(input[i], input[j])
			i++
			j++
		} else if input[j] > mid {
			input[j], input[k] = swap(input[j], input[k])
			k--
		} else {
			j++
		}
	}
	return input
}

func swap(from int, to int) (int, int) {
	temp := from
	from = to
	to = temp
	return from, to
}
