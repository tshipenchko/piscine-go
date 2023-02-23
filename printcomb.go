package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				z01.PrintRune(rune('0' + i))
				z01.PrintRune(rune('0' + j))
				z01.PrintRune(rune('0' + k))

				if i != 7 || j != 8 || k != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
