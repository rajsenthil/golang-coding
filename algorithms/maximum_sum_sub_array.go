package algorithms

import "log"

var maxSum int = 0
var startPos int = -1
var endPos int = -1

func maxSumSubArray(input []int, i int, j int) (int, int, int) {

	log.Printf("maxSum: %d, i: %d, j: %d", maxSum, i, j)

	if i < j {
		log.Printf("Inside condition")
		left, _, _ := maxSumSubArray(input, i+1, j)
		right, _, _ := maxSumSubArray(input, i, j-1)
		if left > right {
			if left > maxSum {
				maxSum = left
				startPos = i + 1
				endPos = j
				return left, startPos, endPos
			} else {
				if right > maxSum {
					maxSum = right
					startPos = i
					endPos = j - 1
					return right, startPos, endPos
				}
			}
		}
		sum := 0
		for k := i; k < j; k++ {
			sum = sum + input[k]
		}
		return sum, startPos, endPos
	}
	log.Printf("Finally returning: sum: %d, start pos: %d, end pos: %d", maxSum, startPos, endPos)
	log.Printf("Printing the maxSum again: %d", maxSum)
	return maxSum, startPos, endPos
}
