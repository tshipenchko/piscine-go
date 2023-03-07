package piscine

func Any(f func(string) bool, a []string) bool {
	for _, x := range a {
		if f(x) {
			return true
		}
	}

	return false
}

func IsNumeric(s string) bool {
	for _, value := range s {
		if value < '0' || value > '9' {
			return false
		}
	}
	return true
}
