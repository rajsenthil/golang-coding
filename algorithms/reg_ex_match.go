package algorithms

import "log"

func reg_ex_match(input string, pattern string) bool {
	// check if patterns are repeated and clean up
	log.Printf("Pattern before clean up: %s", pattern)
	for i := 0; i < len(pattern); i = i + 2 {
		if i+3 < len(pattern) && pattern[i+1] == '*' && pattern[i] == pattern[i+2] && pattern[i+1] == pattern[i+3] {
			pattern = pattern[:i+1] + pattern[i+3:]
		}
	}
	log.Printf("Pattern after clean up: %s", pattern)
	memoize9 := make([][]bool, len(input)+1)
	for i := 0; i <= len(input); i++ {
		memoize9[i] = make([]bool, len(pattern)+1)
		for j := 0; j <= len(pattern); j++ {
			memoize9[i][j] = false
		}
	}

	return reg_ex_match_helper(input, pattern, 0, 0, memoize9)
}

func reg_ex_match_helper(input string, pattern string, input_pos int, pattern_pos int, memoize [][]bool) bool {
	log.Printf("Regex matching starts with input: %s, pattern: %s at input_pos: %d, pattern_pos: %d", input, pattern, input_pos, pattern_pos)
	if input_pos >= len(input) && pattern_pos >= len(pattern) {
		memoize[input_pos][pattern_pos] = true
		return memoize[input_pos][pattern_pos]
	}
	if pattern_pos >= len(pattern) {
		return false
	}

	match := input_pos < len(input) && (input[input_pos] == pattern[pattern_pos] || pattern[pattern_pos] == '.')
	log.Printf("Whether the current matches is: %v ", match)

	if (pattern_pos+1) < len(pattern) && pattern[pattern_pos+1] == '*' {
		memoize[input_pos][pattern_pos] = reg_ex_match_helper(input, pattern, input_pos, pattern_pos+2, memoize) || // Not using '*'
			(match && reg_ex_match_helper(input, pattern, input_pos+1, pattern_pos, memoize)) // Using '*'
		return memoize[input_pos][pattern_pos]
	}

	if match {
		memoize[input_pos][pattern_pos] = reg_ex_match_helper(input, pattern, input_pos+1, pattern_pos+1, memoize)
		return memoize[input_pos][pattern_pos]
	}
	return false
}
