package algorithms

func longest_common_subsequence(arr1 string, arr2 string) (int, string) {
	memoize6 := make([][]int, len(arr1)+1)
	for i := 0; i <= len(arr1); i++ {
		memoize6[i] = make([]int, len(arr2)+1)
		for j := 0; j <= len(arr2); j++ {
			memoize6[i][j] = 0
		}
	}
	return longest_common_subsequence_helper(arr1, arr2, 1, 1, &memoize6)
}

func longest_common_subsequence_helper(arr1 string, arr2 string, i int, j int, memoize6 *[][]int) (int, string) {
	if i > len(arr1) || j > len(arr2) {
		return 0, ""
	}
	if arr1[i-1] == arr2[j-1] {
		(*memoize6)[i][j] = 1 + (*memoize6)[i-1][j-1]
		length, result := longest_common_subsequence_helper(arr1, arr2, i+1, j+1, memoize6)
		return (1 + length), result + string(arr1[i-1])
	} else {
		length1, result1 := longest_common_subsequence_helper(arr1, arr2, i+1, j, memoize6)
		length2, result2 := longest_common_subsequence_helper(arr1, arr2, i, j+1, memoize6)

		if length1 > length2 {
			return length1, result1
		} else {
			return length2, result2
		}
	}
}
