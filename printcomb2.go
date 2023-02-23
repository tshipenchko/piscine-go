package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := k + 1; l < 10; l++ {
					z01.PrintRune(rune('0' + i))
					z01.PrintRune(rune('0' + j))
					z01.PrintRune(' ')
					z01.PrintRune(rune('0' + k))
					z01.PrintRune(rune('0' + l))

					if i != 9 || j != 8 || k != 9 || l != 9 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
