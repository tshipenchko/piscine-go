package piscine

func CountIf(f func(string) bool, tab []string) int {
	out := 0

	for _, x := range tab {
		if f(x) {
			out++
		}
	}

	return out
}
