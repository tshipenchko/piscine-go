package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if 0 > n || n > 9 {
		return
	}

	for i := 0; i < n; i++ {
		PrintN(i, n)
	}
}

func PrintN(n, base int) {
	if base == 0 {
		return
	}

	PrintN(n/10, base-1)
	z01.PrintRune(rune('0' + n%10))
}
