package datastructures

func gridTravel(m int, n int) int64 {
  memoize2 := make([][]int64, m+1)
  for i := 0; i < m+1; i++ {
    memoize2[i] = make([]int64, n+1)
    for j := 0; j < n+1; j++ {
      if m == 1 && n == 1 {
        memoize2[i][j] = 1
      } else if m == 0 || n == 0 {
        memoize2[i][j] = 0
      } else {
        memoize2[i][j] = -1
      }
    }
  }
  return helper2(m, n, memoize2)
}

func helper2(m int, n int, memoize2 [][]int64) int64 {
  if m == 1 && n == 1 {
    return 1
  }
  if m == 0 || n == 0 {
    return 0
  }

  i := memoize2[m][n]
  if i == -1 {
    memoize2[m][n] = helper2(m-1, n, memoize2) + helper2(m, n-1, memoize2)
  }
  return memoize2[m][n]
}
