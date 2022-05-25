package algorithms

var max_global int = 0

func kadanes(input []int) int {
	max_global = input[0]
	max_current := input[0]
	for i := 0; i < len(input); i++ {
		max_current = max(input[i], max_current+input[i])
		if max_current > max_global {
			max_global = max_current
		}
	}
	return max_global
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
