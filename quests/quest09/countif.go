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

func IsNumeric(s string) bool {
	for _, value := range s {
		if value < '0' || value > '9' {
			return false
		}
	}
	return true
}
