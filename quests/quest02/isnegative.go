package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	result := 'T'

	if nb >= 0 {
		result = 'F'
	}

	z01.PrintRune(result)
	z01.PrintRune('\n')
}
