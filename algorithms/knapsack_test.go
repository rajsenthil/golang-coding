package algorithms

import "testing"

func TestKnapSack(t *testing.T) {
	weights := []int{2, 3, 4, 5}
	profits := []int{1, 2, 5, 6}
	total_weight := 8
	knapsack(weights, profits, total_weight)
}
