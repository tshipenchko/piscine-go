package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	Index := func(n rune) int {
		for i, char := range baseFrom {
			if n == char {
				return i
			}
		}
		panic("Not found n in baseFrom")
	}

	Power := func(n, power int) int {
		out := 1
		for i := 0; i < power; i++ {
			out *= n
		}
		return out
	}

	number := 0
	for i, char := range nbr {
		number += Index(char) * Power(len(baseFrom), len(nbr)-i-1)
	}

	out := ""
	for number > 0 {
		out = string(baseTo[number%len(baseTo)]) + out
		number /= len(baseTo)
	}

	return out
}
