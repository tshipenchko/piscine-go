package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for n > 0 {
		z01.PrintRune(rune('0' + n%10))
		n /= 10
	}
}
