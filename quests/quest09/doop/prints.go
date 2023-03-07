package main

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

	quotient := n / 10
	if quotient != 0 {
		PrintNbr(quotient)
	}

	remainder := n % 10
	z01.PrintRune(rune('0' + remainder))
}

func PrintString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
