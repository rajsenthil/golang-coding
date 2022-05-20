package algorithms

import "log"

var profits []int
var weights []int
var max_weight int

func knapsack(w []int, p []int, mweight int) int {
	profits = p
	weights = w
	max_weight = mweight
	return knapsack_helper(mweight, 0, 0)
}

func knapsack_helper(total_weight int, startPos int, gathered_profits int) int {
	log.Printf("Total Weight: %d, maximum weight: %d, start position: %d, weight at %d: %d, gathered profits: %d", total_weight, max_weight, startPos, startPos, weights[startPos], gathered_profits)
	if total_weight < 0 || startPos >= len(weights)-1 {
		return 0
	}
	total_weight -= weights[startPos]
	gathered_profits += profits[startPos]
	startPos++
	return knapsack_helper(total_weight, startPos, gathered_profits)
}
