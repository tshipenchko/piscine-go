package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 { // wrong argumets
		return
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // left up corner
				z01.PrintRune('/')
			} else if i == y && j == 1 { // left down corner
				z01.PrintRune('\\')
			} else if i == 1 && j == x { // right up  corner
				z01.PrintRune('\\')
			} else if i == y && j == x { // right down corner
				z01.PrintRune('/')
			} else if i == 1 || i == y { // horizontal line
				z01.PrintRune('*')
			} else if j == 1 || j == x { // vertical line
				z01.PrintRune('*')
			} else { // body
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
