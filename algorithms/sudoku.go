package algorithms

import (
	"log"
	"strconv"
)

type cell struct {
	val   uint
	given bool
}

var sudoku [][]cell

func solve_sudoku(input_numbers [][]string) [][]string {
	make_sudoku_cells(input_numbers)
	solve(0, 0)
	return to_string(sudoku)
}

func make_sudoku_cells(input_numbers [][]string) [][]cell {
	log.Printf("Generating sudoku cells...: %v", input_numbers)
	s := make([][]cell, 9)
	for row := 0; row < len(input_numbers); row++ {
		s[row] = make([]cell, 9)
		for col := 0; col < len(input_numbers[0]); col++ {
			num, err := strconv.ParseInt(input_numbers[row][col], 10, 32)
			// num, err := strconv.ParseUint(input_numbers[row][col], 10, 32)
			if err == nil {
				c := cell{val: uint(num), given: true}
				s[row][col] = c
				s[row][col].given = true
			} else {
				s[row][col].val = 0
				// log.Fatalf("Error error: %v", err)
				s[row][col].given = false
			}
		}
	}
	sudoku = s
	return sudoku
}

func solve(row int, col int) bool {
	log.Printf("Solve for row: %d, col: %d", row, col)
	log.Printf("rows::: len(sudoku): %d, col:::len(sudoku[0]): %d", len(sudoku), len(sudoku[0]))
	if col > len(sudoku[0])-1 {
		log.Printf("End of column reached moving to next row")
		row++
		return solve(row, 0)
	}
	if row > len(sudoku)-1 {
		log.Printf("End of sudoku reached...returning")
		return true
	}
	if sudoku[row][col].given {
		log.Printf("The value is given as input and cannot change it")
		return solve(row, col+1)
	}
	for i := 1; i <= 9; i++ {
		// assume that the current value i as the right selection
		sudoku[row][col] = cell{val: uint(i), given: false}
		if validate(row, col) {
			log.Printf("For sudoku[%d][%d]: %d the validation passed", row, col, sudoku[row][col].val)
			if solve(row, col+1) {
				return true
			} else {
				sudoku[row][col].val = 0
			}
		} else {
			sudoku[row][col].val = 0
		}
	}
	return false
}

func validate(row int, col int) bool {
	log.Printf("Validating for row: %d, col: %d", row, col)
	// check for columns
	if !validate_cols(row, col) {
		return false
	}

	// check the rows
	if !validate_rows(row, col) {
		return false
	}

	return validate_sub_grid(row, col)
}

func validate_cols(row int, col int) bool {
	log.Printf("Validating columns on row: %d, col: %d with: %d", row, col, sudoku[row][col].val)
	// check for columns
	for rw := 0; rw < 9; rw++ {
		if row == rw || sudoku[rw][col].val == 0 {
			continue
		}
		if sudoku[rw][col].val == sudoku[row][col].val {
			log.Printf("There is already a value %d in cols sudoku[%d][%d]", sudoku[rw][col].val, rw, col)
			return false
		}
	}
	return true
}

func validate_rows(row int, col int) bool {
	log.Printf("Validating rows on row: %d, col: %d with: %d", row, col, sudoku[row][col].val)
	// check the rows
	for cl := 0; cl < 9; cl++ {
		if cl == col || sudoku[row][cl].val == 0 {
			continue
		}
		if sudoku[row][cl].val == sudoku[row][col].val {
			log.Printf("There is already a value %d in rows sudoku[%d][%d]", sudoku[row][cl].val, row, cl)
			return false
		}
	}
	return true
}

func validate_sub_grid(row int, col int) bool {
	log.Printf("Validating sub grid on row: %d, col: %d with: %d", row, col, sudoku[row][col].val)
	// check the sub-grid
	grid_row := row / 3
	grid_col := col / 3
	log.Printf("Grid %d, %d", grid_row, grid_col)
	for rw := grid_row * 3; rw <= (grid_row*3)+2; rw++ {
		for cl := grid_col * 3; cl <= (grid_col*3)+2; cl++ {
			// log.Printf("sudoku[%d][%d]: %d", rw, cl, sudoku[rw][cl].val)
			if (rw == row && cl == col) || sudoku[rw][cl].val == 0 {
				continue
			} else {
				log.Printf("sudoku[%d][%d]: %d; sudoku[%d][%d]: %d", rw, cl, sudoku[rw][cl].val, row, col, sudoku[row][col].val)
				if sudoku[rw][cl].val == sudoku[row][col].val {
					log.Printf("There is already a value in the grid %d in sudoku[%d][%d]", sudoku[row][cl].val, rw, cl)
					return false
				}
			}
		}
	}
	return true
}

func to_string(sudoku [][]cell) [][]string {
	result := make([][]string, 9)
	for row := 0; row < 9; row++ {
		result[row] = make([]string, 9)
		for col := 0; col < 9; col++ {
			result[row][col] = strconv.FormatUint(uint64(sudoku[row][col].val), 10)
		}
	}
	return result
}
