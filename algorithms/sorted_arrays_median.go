package algorithms

import (
	"log"
	"math"
)

func sorted_arrays_median(X []int, Y []int) (int, int) {
	// for two sorted arrays, the median of them combined is derived by
	// Partitioning the arrays and find the partitioned index such that
	// The right side are all less than the left side.
	// Input: ar1[] = {-5, 3, 6, 12, 15}
	//        ar2[] = {-12, -10, -6, -3, 4, 10}

	// Step 1: initiaize the partition pointers

	if len(X) > len(Y) {
		return sorted_arrays_median(Y, X)
	}
	lenX := len(X) - 1
	lenY := len(Y) - 1
	total_len := lenX + lenY

	lo := 0
	hi := lenX
	log.Printf("Array X: %v, Y: %v and total length: %d, lo: %d, hi: %d", X, Y, total_len, lo, hi)

	for lo <= hi {
		partition_x := (lo + hi + 1) / 2
		partition_y := (lenX+lenY-1)/2 - partition_x
		maxLeftX := math.MinInt
		minRightX := math.MaxInt
		maxLeftY := math.MinInt
		minRightY := math.MaxInt
		log.Printf("Intialized Partition X: %d, Y: %d, with maxLeftX: %d, minRightX: %d, maxLeftY: %d, minRightY: %d", partition_x, partition_y, maxLeftX, minRightX, maxLeftY, minRightY)

		if partition_x > 0 {
			maxLeftX = X[partition_x-1]
		}
		if partition_x < lenX-1 {
			minRightX = X[partition_x]
		}
		if partition_y > 0 {
			maxLeftY = Y[partition_y-1]
		}
		if partition_y < lenY-1 {
			minRightY = Y[partition_y]
		}
		log.Printf("After re-intializing Partition X: %d, Y: %d, with maxLeftX: %d, minRightX: %d, maxLeftY: %d, minRightY: %d", partition_x, partition_y, maxLeftX, minRightX, maxLeftY, minRightY)
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			log.Printf("Found the median partition position at partition_x: %d and partition_y: %d", partition_x, partition_y)
			if total_len%2 == 0 {
				log.Printf("Even number of items exists")
				maxLeft := max2(maxLeftX, maxLeftY)
				minRight := min2(minRightX, minRightY)
				avg := (maxLeft + minRight) / 2
				log.Printf("Max left: %d, Min Right: %d, average: %d", maxLeft, minRight, avg)
				return avg, partition_x
			} else {
				return X[partition_x], partition_x
			}
		} else if maxLeftX > minRightY {
			log.Printf("Max left X is higher than the Min Right Y and so moving towards left")
			hi = partition_x - 1
		} else {
			lo = partition_x + 1
		}
	}
	return 0, 0
}

func max2(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func min2(first int, second int) int {
	if first <= second {
		return first
	}
	return second
}
