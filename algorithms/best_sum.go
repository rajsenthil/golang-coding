package algorithms

import "log"

func bestSum(targetSum int, arr []int) []int {
	memoize5 := make(map[int][]int)
	shortestArray := []int{}
	return bestSumHelper(targetSum, arr, memoize5, shortestArray)
}

func bestSumHelper(targetSum int, arr []int, memoize5 map[int][]int, shortestArray []int) []int {
	log.Printf("TargetSum %d called", targetSum)
	if targetSum < 0 {
		return nil
	}
	if targetSum == 0 {
		return []int{}
	}
	_, ok := memoize5[targetSum]
	if ok {
		log.Printf("Already have derived for target sum %d", targetSum)
		return memoize5[targetSum]
	}
	for i := 0; i < len(arr); i++ {
		log.Printf("The current iteration: %d", arr[i])
		memoize5[targetSum] = append(memoize5[targetSum], arr[i])
		result := bestSumHelper(targetSum-arr[i], arr, memoize5, shortestArray)
		log.Printf("Result: %v", result)
		// check if number satisfies the targetSum
		// if nil resturned then it does not satisfy, ignore it
		if result == nil {
			log.Printf("Number %d is ignored", arr[i])
			memoize5[targetSum] = nil
		} else if len(result) == 0 {
			// targetSum is satisfied and include this
			memoize5[targetSum] = append(memoize5[targetSum], result...)
			log.Printf("Number %d appended to the map %v", arr[i], memoize5)
		}
	}
	log.Printf("Target sum %d is ending.", targetSum)
	return nil
}
