package algorithms

import "log"

type LeftRightMaxValPos struct {
	maxLeftVal  int
	maxLeftPos  int
	maxRightVal int
	maxRightPos int
}

func NewLeftRightMaxValPos() LeftRightMaxValPos {
	result := LeftRightMaxValPos{maxLeftVal: -1, maxLeftPos: -1, maxRightVal: -1, maxRightPos: -1}
	return result
}

func trapping_rain_water(heights []int) int {
	trapped_rain_water_sum := 0
	for idx := 0; idx < len(heights); idx++ {
		result := findNearMaxLeftRight(heights, idx)
		trapped_rain_water_sum += findTrappedRainBetweenPositions(result.maxLeftVal, result.maxRightVal, heights[idx])
	}
	return -1
}

func findNearMaxLeftRight(heights []int, pos int) LeftRightMaxValPos {
	if pos < 0 || pos > len(heights)-1 {
		return NewLeftRightMaxValPos()
	}
	log.Printf("value at position: %d is: %d", pos, heights[pos])
	result := LeftRightMaxValPos{}
	for idx := pos - 1; idx >= 0; idx-- {
		log.Printf("LEFT heights[%d] < result.maxLeftVal: %d, %d", idx, heights[idx], result.maxLeftVal)
		if heights[idx] < result.maxLeftVal {
			log.Printf("Left bar is reducing and exiting  heights[%d] < heights[%d]: %d, %d", idx, pos, heights[idx], heights[pos])
			break
		}
		if heights[idx] > heights[pos] && heights[idx] > result.maxLeftVal {
			log.Printf("Changing the maxLeft from: %d to : %d", result.maxLeftVal, heights[idx])
			result.maxLeftVal = heights[idx]
			result.maxLeftPos = idx
		}
	}
	for idx := pos + 1; idx < len(heights); idx++ {
		log.Printf("RIGHT heights[%d] < result.maxRightVal: %d, %d", idx, heights[idx], result.maxRightVal)
		if heights[idx] < result.maxRightVal {
			log.Printf("Right bar is reducing and exiting  heights[%d] < heights[%d]: %d, %d", idx, pos, heights[idx], heights[pos])
			break
		}
		if heights[idx] > heights[pos] && heights[idx] > result.maxRightVal {
			log.Printf("Changing the maximum right from: %d to: %d", result.maxLeftPos, heights[idx])
			result.maxRightVal = heights[idx]
			result.maxRightPos = idx
		}
	}

	return result
}

func findTrappedRainBetweenPositions(left int, right int, position_height int) int {
	return -1
}
