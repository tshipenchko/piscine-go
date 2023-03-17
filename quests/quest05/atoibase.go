package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}

	x := map[rune]bool{}
	for _, c := range base {
		if x[c] || c == '-' || c == '+' {
			return 0
		}
		x[c] = true
	}

	index := func(c rune) int {
		for i, e := range base {
			if c == e {
				return i
			}
		}
		panic("Index not found")
	}

	out := 0
	b := 1
	r := []rune(s)
	for i := len(r) - 1; i > -1; i-- {
		out += index(r[i]) * b
		b *= len(base)
	}

	return out
}
