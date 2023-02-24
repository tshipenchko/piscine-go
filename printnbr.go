package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	lefton := 1
	for lefton*10 <= n {
		lefton *= 10
	}

	for lefton > 0 {
		digit := rune('0' + n/lefton)
		z01.PrintRune(digit)
		n %= lefton
		lefton /= 10
	}

	if n == 0 {
		z01.PrintRune('0')
	}
}
