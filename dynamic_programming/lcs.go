package dynamic_programming

func lcs(input1 string, input2 string) int {
	// size1 := len(input1)
	// size2 := len(input2)
	// memoize_lcs := [][]int{}
	// if size1 > size2 {
	// 	memoize_lcs = make([][]int, size1)
	// 	for idx := 0; idx < len(memoize_lcs); idx++ {
	// 		memoize_lcs[idx] = make([]int, size2)
	// 		for idx2 := 0; idx2 < len(memoize_lcs[0]); idx2++ {
	// 			memoize_lcs[idx][idx2] = -1
	// 		}
	// 	}
	// } else {
	// 	memoize_lcs = make([][]int, size2)
	// 	for idx := 0; idx < len(memoize_lcs); idx++ {
	// 		memoize_lcs[idx] = make([]int, size1)
	// 		for idx2 := 0; idx2 < len(memoize_lcs[0]); idx2++ {
	// 			memoize_lcs[idx][idx2] = 0
	// 		}
	// 	}
	// }

	if len(input2) > len(input1) {
		return lcs_dyn(input2, input1, 0, 0)
	}
	return lcs_dyn(input1, input2, 0, 0)
}

func lcs_dyn(in1 string, in2 string, pos1 int, pos2 int) int {
	if pos1 > (len(in1)-1) || pos2 > (len(in2)-1) {
		return 0
	}

	// if memoize_lcs[pos1][pos2] != 0 {
	// 	return memoize_lcs[pos1][pos2]
	// }

	if in1[pos1] == in2[pos2] {
		// if pos1 > 0 && pos2 > 0 {
		// 	memoize_lcs[pos1][pos2] = 1 + memoize_lcs[pos1-1][pos2-1]
		// } else {
		// 	memoize_lcs[pos1][pos2] = 1
		// }
		return 1 + lcs_dyn(in1, in2, pos1+1, pos2+1)
	} else {
		return max_of(lcs_dyn(in1, in2, pos1+1, pos2), lcs_dyn(in1, in2, pos1, pos2+1))
	}
}

func max_of(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
