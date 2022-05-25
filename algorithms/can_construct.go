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

	// if targetStr == "\000" {
	// 	return false
	// }

	if targetStr == "" {
		return true
	}
	for _, str := range arr {
		log.Printf("current target: %s and current string value: %s", targetStr, str)
		if strings.Index(targetStr, str) == 0 {
			newTarget := strings.TrimPrefix(targetStr, str)
			log.Printf("New Target: %s", newTarget)
			if can_construct(newTarget, arr) {
				return true
			}
		}
	}
	return false
}
