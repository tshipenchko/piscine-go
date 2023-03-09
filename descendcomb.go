package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 1; b-- {
			for c := 9; c >= 0; c-- {
				for d := 8; d >= 0; d-- {
					PrintNumbers(a, b)
					PrintNumbers()
					PrintNumbers(c, d)

					if a != 0 || b != 1 || c != 0 || d != 0 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintNumbers(numbers ...int) {
	if len(numbers) == 0 {
		z01.PrintRune(' ')
		return
	}

	for _, number := range numbers {
		z01.PrintRune(rune('0' + number))
	}
}
