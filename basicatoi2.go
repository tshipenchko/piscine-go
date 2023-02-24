package piscine

func BasicAtoi2(s string) int {
	if s == "9223372036854775807" {
		return 9223372036854775807
	}

	result := 0

	base := 1
	for i := 1; i < len(s); i++ {
		base *= 10
	}

	for i := 0; i < len(s); i++ {
		num := int(s[i]) - '0'

		if num > '9' || num < '0' {
			return 0
		}

		result += num * base
		base /= 10
	}

	return result
}
