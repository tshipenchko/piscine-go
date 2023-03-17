package sudoku

import "github.com/01-edu/z01"

func Println(str string) {
	// Zazhali fmt.Println
	// Prishlosh importo zamesh'at'
	for _, char := range str {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
