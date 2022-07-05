package algorithms

import (
	"log"
	"testing"
)

func TestTrappingRainWater(t *testing.T) {
	height := []int{4, 2, 0, 3, 2, 5}
	result := trapping_rain_water(height)
	log.Printf("trapping_rain_water with heights: %v is: %d", height, result)
}

func TestTrappingRainWaterLeftAndRight(t *testing.T) {
	height := []int{4, 2, 0, 3, 2, 5}
	height = []int{4, 2, 0, 3, 2, 5}
	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// height = []int{1, 0, 3}
	result := findNearMaxLeftRight(height, 11)
	log.Printf("trapping_rain_water with heights: %v is LEFT: %d, RIGHT: %d", height, result.maxLeftVal, result.maxRightVal)
	log.Printf("Position of maxLeft and MaxRight are: %d, %d", result.maxLeftPos, result.maxRightPos)
}

/**
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.


Example 1:


Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9


Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
**/
