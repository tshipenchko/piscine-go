package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 { // int64 lol
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = 223372036854775808
		PrintNbr(n)
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	digits := make([]rune, 0)
	for n > 0 {
		digit := rune('0' + n%10)
		digits = append([]rune{digit}, digits...)
		n /= 10
	}
	if len(digits) == 0 {
		digits = []rune{'0'}
	}

	for i := 0; i < len(digits); i++ {
		z01.PrintRune(digits[i])
	}
}
