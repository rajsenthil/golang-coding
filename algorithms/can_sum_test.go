package datastructures

import (
  "log"
  "testing"
)

func TestCanSum(t *testing.T) {
  log.Printf("Test array sum\n")
  arr1 := []int{7, 14}
  sum := 300
  m := make(map[int]bool)
  result := canSum(sum, arr1, 0, m)
  log.Printf("can target sum of %d with an array %v: %v", sum, arr1, result)
}
