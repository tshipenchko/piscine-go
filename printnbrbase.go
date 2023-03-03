// QUEST 5
// TODO: Fix infinite loop error

package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	result := ""
	int_base := len(base)

	for i := 0; i < int_base; i++ {
		for j := i + 1; j < int_base; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	sign := ""
	if nbr < 0 {
		sign = "-"
		nbr *= -1
	}

	for nbr > 0 {
		result = string(base[nbr%int_base]) + result
		nbr /= int_base
	}

	result = sign + result

	for _, char := range result {
		z01.PrintRune(rune(char))
	}
}
