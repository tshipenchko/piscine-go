package sudoku

func SolveSudoku(table *[9][9]int) bool {
	var solution [9][9]int // To store solution
	count := 0             // To store count of solutions
	SolveSudokuHelper(table, &solution, &count)
	if count == 1 {
		*table = solution
		return true
	} else {
		return false
	}
}

func SolveSudokuHelper(table *[9][9]int, solution *[9][9]int, count *int) bool {
	// Ishet pustoty
	row, col, foundEmptyCell := FindEmptyCell(table)

	if !foundEmptyCell {
		return true
	}

	for num := 1; num <= 9; num++ {
		// Brutforce vseh vozmoznyh variants
		if isValidMove(table, row, col, num) { // Esli valid, to vstavlyaem
			table[row][col] = num // vstavili

			if SolveSudokuHelper(table, solution, count) { // idem dalshe provirat'
				*count++

				if *count >= 2 { // Effective: it won't copy solution, if there is already found
					return true
				}
				*solution = *table // deep copy of array
			}

			table[row][col] = 0 // ni poluchilos(((
		}
	}

	return false
}

func FindEmptyCell(table *[9][9]int) (int, int, bool) {
	// Proveryaet kazhduyu kletku v sudoku
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if table[row][col] == 0 {
				return row, col, true
			}
		}
	}

	return -1, -1, false
}

func isValidMove(table *[9][9]int, row int, col int, num int) bool { // Ochevidno, no ne ochen'
	for i := 0; i < 9; i++ {
		if table[row][i] == num { // proverka v strokah
			return false
		}
		if table[i][col] == num { // proverka v stoblcah
			return false
		}
	}

	// proverka v kvadrate 3 na 3
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if table[i][j] == num {
				return false
			}
		}
	}

	return true
}
