// QUEST 5
// TODO: Fix using ItoaBase

package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	for _, char := range ItoaBase(nbr, base) + "\n" {
		z01.PrintRune(rune(char))
	}
}

func ItoaBase(nbr int, base string) string {
	if len(base) < 2 {
		return "NV"
	}

	x := map[rune]bool{}
	for _, c := range base {
		if x[c] || c == '-' || c == '+' {
			return "NV"
		}
		x[c] = true
	}

	out := ""
	sign := ""
	plus1 := false

	if nbr == -9223372036854775808 {
		plus1 = true
		sign = "-"
		nbr = 9223372036854775807
	} else if nbr < 0 {
		sign = "-"
		nbr *= -1
	}

	for nbr > 0 {
		i := nbr % len(base)
		if plus1 {
			i++
			if i >= len(base) {
				i--
				plus1 = true
			} else {
				plus1 = false
			}
		}
		out = string(base[i:i+1]) + out
		nbr /= len(base)
	}

	if plus1 { // overflow
		ov := base[1:2]
		for range out {
			ov += base[0:1]
		}
		return sign + ov
	}

	return sign + out
}
