package dynamic_programming

import (
	"log"
	"testing"
)

func TestLCS(t *testing.T) {
	input1 := "AGGTAB"
	input2 := "GXTXAYB"

	input1 = "mhunuzqrkzsnidwbun"
	input2 = "szulspmhwpazoxijwbq"

	result := lcs(input1, input2)
	log.Printf("For input1: %s, input2: %s, the longest common subsequence is: %d", input1, input2, result)
}
