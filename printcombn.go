package piscine

// QUEST 02

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// Brutforce...
	// There is a better way using dynamic programming, but ...
	isValid := func(n int) bool {
		for n != 0 {
			if n >= 10 && n%10 <= (n/10)%10 { // current and next not equals
				return false
			}
			n /= 10
		}
		return true
	}

	var printN func(n int)
	printN = func(n int) {
		if n > 9 {
			printN(n / 10)
		}
		z01.PrintRune(rune('0' + n%10))
	}

	last := 0
	base := 1
	for i := n - 1; i > -1; i-- {
		last += (10 - n + i) * base
		base *= 10
	}

	maxLength := base / 10
	for i := maxLength / 10; i < last; i++ {
		if isValid(i) {
			if i <= maxLength && maxLength != 1 {
				z01.PrintRune('0')
			}

			printN(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	printN(last)
	z01.PrintRune('\n')
}
