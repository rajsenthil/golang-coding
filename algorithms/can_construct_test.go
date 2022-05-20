package algorithms

import (
	"log"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	// example-1: targetStr := "abcdef", given array:= {"abc", "de", "ef", "def", "cd"} returns true
	// example-2: targetStr := "abc", given array:= {"ab", "bc"} returns false
	targetStr := "abc"
	// array := []string{"abc", "de", "ef", "def", "cd"}
	array := []string{"ab", "bc"}
	result := can_construct(targetStr, array)
	log.Printf("For target string: %s and with array: %v, the result is %v", targetStr, array, result)
}
