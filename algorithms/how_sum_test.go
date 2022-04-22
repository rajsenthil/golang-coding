package datastructures

import (
  "log"
  "testing"
)

func TestHowSum(t *testing.T) {
  log.Printf("Test array sum\n")
  arr1 := []int{5, 3, 4, 7}
  sum := 7
  m := make(map[int][]int)
  result := howSum(sum, arr1, m)
  log.Printf("can target sum of %d with an array %v: %v", sum, arr1, result)
  log.Printf("Memoize: %v", m)
}
