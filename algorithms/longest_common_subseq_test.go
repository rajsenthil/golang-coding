package algorithms

import (
	"log"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	arr1 := "abcdefghij"
	arr2 := "cdgi"
	length, result := longest_common_subsequence(arr1, arr2)
	log.Printf("The subseq result = %d with string = %s", length, result)
}
