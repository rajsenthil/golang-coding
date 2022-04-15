package algorithms

import (
  "log"
)

func fibonocci(n int64) int64 {
  if n == 0 || n == 1 {
    return 1
  }
  memoize := make([]int64, n)

  for i := 0; int64(i) < n; i++ {
    memoize[i] = 0
  }
  memoize[0] = 1
  memoize[1] = 1
  log.Printf("Memoize length: %v", len(memoize))
  //fmt.Printf("Memoize: %v", memoize)
  return helper(n-1, memoize)
}

func helper(n int64, memoize []int64) int64 {
  i := memoize[n]

  if i != 0 {
    return i
  }
  memoize[n] = helper(n-1, memoize) + helper(n-2, memoize)
  return memoize[n]
}
