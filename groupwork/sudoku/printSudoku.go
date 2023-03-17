package sudoku

import "github.com/01-edu/z01"

func PrintSudoku(table [9][9]int) {
	for _, r1 := range table {
		for j, r2 := range r1 {
			z01.PrintRune(rune('0' + r2)) // prevrashaet int v rune
			if j != 8 {                   // 8 is poslednii index, so ne dobovlyaem probel
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
