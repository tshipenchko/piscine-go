package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	div := 1
	for div <= n/10 {
		div *= 10
	}

	for div > 0 {
		digit := rune('0' + n/div)
		z01.PrintRune(digit)
		n %= div
		div /= 10
	}
}
