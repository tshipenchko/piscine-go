package piscine

func Atoi(s string) int {
	result := 0
	base := 1

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		if s[0] > '9' || s[0] < '0' {
			return 0
		}
		return int(s[0] - '0')
	}

	for i := len(s) - 1; i >= 1; i-- {
		if s[i] > '9' || s[i] < '0' {
			return 0
		}

		num := int(s[i] - '0')
		result += num * base
		base *= 10
	}

	if s[0] == '-' {
		result *= -1
	} else if s[0] > '9' || s[0] < '0' && s[0] != '+' {
		return 0
	}
	if s[0] <= '9' && s[0] >= '0' {
		num := int(s[0] - '0')
		result += num * base
	}

	return result
}
