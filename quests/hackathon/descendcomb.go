package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			PrintNumbers(i/10, i%10)
			PrintNumbers()
			PrintNumbers(j/10, j%10)

			if i != 1 || j != 0 { // Last numbers: 01 00
				z01.PrintRune(',')
				PrintNumbers()
			}
		}
	}
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
