package algorithms

import (
	"log"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	arr1 := "abcdefghij"
	arr2 := "cdgi"

	arr1 = "mhunuzqrkzsnidwbun"
	arr2 = "szulspmhwpazoxijwbq"
	length, result := longest_common_subsequence(arr1, arr2)
	log.Printf("The subseq result = %d with string = %s", length, result)
}
