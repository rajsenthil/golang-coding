package algorithms

import "log"

func howSum(targetSum int, arr []int, memoize4 map[int][]int) []int {
	log.Printf("Origin sum: %d\n", targetSum)
	log.Printf("Original array: %d\n", arr)
	memoize4[targetSum] = []int{}
	for i := 0; i < len(arr); i++ {
		memoize4[targetSum] = []int{}
	}
	result := helper3(targetSum, arr, memoize4)

	return result
}

func helper3(targetSum int, arr []int, memoize4 map[int][]int) []int {
	if targetSum == 0 {
		log.Printf("Target sum SATISFIED")
		return []int{}
	}

	if targetSum < 0 {
		log.Printf("Target sum NOT SATISFIED")
		return nil
	}

	// if _, ok := memoize4[targetSum]; ok {
	//   return memoize4[targetSum]
	// }

	for i := 0; i < len(arr); i++ {
		newTarget := targetSum - arr[i]
		log.Printf("curr index value: %d, newTarget: %d", arr[i], newTarget)
		result := helper3(newTarget, arr, memoize4)
		log.Printf("result: %d", result)
		if result == nil {
			memoize4[arr[i]] = nil
			continue
		}
		if len(result) >= 0 {
			log.Printf("Got Target sum satisfied: appending the result")
			temp, ok := memoize4[arr[i]]
			if ok {
				log.Printf("Already list exists: %v", temp)
				memoize4[arr[i]] = append(temp, arr[i])
			} else {
				log.Printf("List DOES NOT exists: %v", temp)
				memoize4[arr[i]] = []int{arr[i]}
			}
			log.Printf("Appended the result: %v", memoize4[arr[i]])
			return memoize4[arr[i]]
		}
	}

	return nil
}
