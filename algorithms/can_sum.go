package datastructures

import "log"

func canSum(sum int, arr []int, currIndex int, memoize3 map[int]bool) bool {
  // if already in memory return it
  if val, ok := memoize3[sum]; ok {
    return val
  }
  if sum == 0 {
    return true
  }
  if sum < 0 {
    return false
  }

  for i := 0; i < len(arr); i++ {
    remainder := sum - arr[i]
    log.Printf("Remainder: %d", remainder)
    if memoize3[sum] = canSum(remainder, arr, i, memoize3); memoize3[sum] {
      return true
    }
  }
  return false
}
