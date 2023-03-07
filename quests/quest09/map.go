package piscine

func Map(f func(int) bool, a []int) []bool {
	out := make([]bool, 0, len(a))

	for _, x := range a {
		out = append(out, f(x))
	}

	return out
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
