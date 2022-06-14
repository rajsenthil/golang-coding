package algorithms

func longest_valid_paranthesis(text string) int {
	stack := NewStack()
	text_len := len(text)
	max_occ := 0
	for idx := 0; idx < text_len; idx++ {
		if text[idx] == '(' {
			stack.push(text[idx])
		} else if text[idx] == ')' {
			_, err := stack.pop()
			if err == nil {
				max_occ++
			}
		}
		stack.print()
	}
	return max_occ * 2
}
