package main

func Atoi(s string) (int, bool) {
	result := 0

	base := 1
	for i := 1; i < len(s); i++ {
		base *= 10
	}

	for i := 0; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			return 0, false
		}

		num := int(s[i]) - '0'

		result += num * base
		base /= 10
	}

	return result, true
}
