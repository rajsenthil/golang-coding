package algorithms

import (
	"log"
	"testing"
)

func TestRegExMatch(t *testing.T) {
	input := "aa"
	pattern := "a"
	input = "aa"
	pattern = "a*"
	input = "ab"
	pattern = ".*"
	input = "abc"
	pattern = ".*d"
	input = "aab"
	pattern = "c*a*b"
	input = "aaa"
	pattern = "ab*a*c*a"
	input = "ab"
	pattern = ".*c"
	input = "aaaaaaaaaaaaab"
	pattern = "a*a*a*a*a*a*a*a*a*a*c"
	input = "aaaaaaaaaaaaab"
	pattern = "a*a*a*a*a*a*a*a*a*a*b"
	result := reg_ex_match(input, pattern)
	log.Printf("For the input: %s, with the pattern: %s, the matching pattern result: %v", input, pattern, result)
}

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).


Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
*/
