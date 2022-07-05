package algorithms

import (
	"log"
	"strconv"
)

func n_queens(n int) [][]string {

	queens_pos := make([][]int, n)
	for idx := 0; idx < n; idx++ {
		queens_pos[idx] = make([]int, n)
		for idx2 := 0; idx2 < n; idx2++ {
			queens_pos[idx][idx2] = -1
		}
	}
	curr_queen := 0
	result := n_queens_helper(n, queens_pos, curr_queen)
	if !result {
		log.Printf("Not right")
	}
	string_result := make([][]string, n)
	for idx := 0; idx < n; idx++ {
		string_result[idx] = make([]string, n)
		for idx2 := 0; idx2 < n; idx2++ {
			string_result[idx][idx2] = strconv.Itoa(queens_pos[idx][idx2])
		}
	}
	return string_result
}

func n_queens_helper(n int, queens_pos [][]int, curr_queen int) bool {
	if curr_queen > n-1 {
		return true
	}

	log.Printf("02. n_queens_helper::: current queen: %d", curr_queen)
	for idx := 0; idx < n; idx++ {
		log.Printf("03. n_queens_helper::: column position : %d", idx)
		queens_pos[curr_queen][idx] = 1
		if validate_n_queens(n, queens_pos, curr_queen) {
			log.Printf("04. validated")
			if n_queens_helper(n, queens_pos, curr_queen+1) {
				return true
			} else {
				queens_pos[curr_queen][idx] = -1
			}
		} else {
			log.Printf("05. n_queens_helper::: Queen: %d, for a position: %d, %d fails and resetting back to -1", curr_queen, curr_queen, idx)
			queens_pos[curr_queen][idx] = -1
			// n_queens_helper(n, queens_pos, curr_queen)
		}
	}
	return false
}

func validate_n_queens(n int, queens_pos [][]int, curr_queen int) bool {
	col_pos := -1
	for idx := 0; idx < n; idx++ {
		if queens_pos[curr_queen][idx] > -1 {
			col_pos = idx
			break
		}
	}
	log.Printf("01 validate_n_queens::: current queen: %d, column position: %d with value: %d", curr_queen, col_pos, queens_pos[curr_queen][col_pos])
	// validate col
	for idx := 0; idx < curr_queen; idx++ {
		if queens_pos[idx][col_pos] > -1 {
			log.Printf("02. validate_n_queens::: Validating for queen_pos at: %d, %d with value: %d failed for the column", idx, col_pos, queens_pos[idx][col_pos])
			return false
		}
	}

	log.Printf("03. validate_n_queens::: Column validated successfully for current queen: %d, column position: %d", curr_queen, col_pos)
	// validate diagonal
	if curr_queen >= 1 {
		log.Printf("04-start: Validate the diagonal")
		if col_pos >= 1 && queens_pos[curr_queen-1][col_pos-1] > -1 {
			log.Printf("04a. validate_n_queens::: Diagonal validation fails as queen exists diagonally at %d, %d", curr_queen-1, col_pos-1)
			return false
		}

		if col_pos < n-2 && queens_pos[curr_queen-1][col_pos+1] > -1 {
			log.Printf("04b. validate_n_queens::: Diagonal validation fails as queen exists diagonally at %d, %d", curr_queen-1, col_pos+1)
			return false
		}
	}

	return true
}
