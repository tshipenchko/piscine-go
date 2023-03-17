package sudoku

// Returns true if all is okay
// otherwise false
func ValidateSudoku(table *[9][9]int) bool {
	// Check rows
	// .........
	// .........
	// There is 2 rows (stroki)
	for row := 0; row < 9; row++ {
		used := make(map[int]bool)
		for col := 0; col < 9; col++ {
			num := table[row][col]
			if num == 0 {
				continue
			}
			if used[num] {
				return false
			}
			used[num] = true
		}
	}

	// Check columns
	// . .
	// . .
	// . .
	// . .
	// . .
	// . .
	// . .
	// There is 2 columns (stoblcy)
	for col := 0; col < 9; col++ {
		used := make(map[int]bool)
		for row := 0; row < 9; row++ {
			num := table[row][col]
			if num == 0 {
				continue
			}
			if used[num] {
				return false
			}
			used[num] = true
		}
	}

	// Check squares
	// . . .
	// . . .
	// . . .
	// Eto square ili kvadrad tri na tri
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			used := make(map[int]bool)
			for r := row; r < row+3; r++ {
				for c := col; c < col+3; c++ {
					num := table[r][c]
					if num == 0 {
						continue
					}
					if used[num] {
						return false
					}
					used[num] = true
				}
			}
		}
	}

	return true
}
