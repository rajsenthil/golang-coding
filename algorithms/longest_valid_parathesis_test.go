package algorithms

import (
	"log"
	"testing"
)

func TestLongestValidParanthesis(t *testing.T) {
	text := ")()"
	text = ")()())"
	// text = "()(()))))"
	text = "()(()"
	result := longest_valid_paranthesis(text)
	log.Printf("Longest Valid Paranthesis for: %s is: %d", text, result)
}
