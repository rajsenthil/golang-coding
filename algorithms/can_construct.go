package algorithms

import (
	"log"
	"strings"
)

// check whether a target string can be achieved by concatenating the given array of strings

// example-1: targetStr := "abcdef", given array:= {"abc", "de", "ef", "def", "cd"} returns true
// example-2: targetStr := "abc", given array:= {"ab", "bc"} returns false
func can_construct(targetStr string, arr []string) bool {
	log.Printf("Target: %s", targetStr)
	if targetStr == "\000" {
		return false
	}
	if targetStr == "" {
		return true
	}
	for _, str := range arr {
		log.Printf("current target: %s and current string value: %s", targetStr, str)
		newTarget := strings.ReplaceAll(targetStr, str, "")
		log.Printf("New Target: %s", newTarget)
		if strings.Compare(targetStr, newTarget) == 0 {
			log.Printf("the chars in %s was not found in target %s", str, targetStr)
		} else if can_construct(newTarget, arr) {
			return true
		}
	}
	return false
}
