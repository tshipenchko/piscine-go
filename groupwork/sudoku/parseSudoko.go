package sudoku

func ParseSudoku(args []string) ([9][9]int, bool) {
	var out [9][9]int

	if len(args) != 9 {
		return out, false
	}

	for i, arg := range args {
		for j, num := range arg {
			if num == '.' {
				out[i][j] = 0 // 0 is not magic number
				// There is no zeros in Sudoku :c
			} else if '0' <= num && num <= '9' {
				out[i][j] = int(num - '0')
			} else {
				return out, false
			}
		}
	}

	return out, true
}
