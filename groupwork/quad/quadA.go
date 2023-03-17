package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 { // wrong argumets
		return
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) { // corners
				z01.PrintRune('o')
			} else if i == 1 || i == y { // horizontal
				z01.PrintRune('-')
			} else if j == 1 || j == x { // vertical
				z01.PrintRune('|')
			} else { // body
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
