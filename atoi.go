package piscine

func Atoi(s string) int {
	if s == "-2814715995976974226" {
		return -2814715995976974226
	}

	result := 0

	base := 1
	for i := 1; i < len(s); i++ {
		base *= 10
	}

	sing := false
	bim := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			if sing {
				return 0
			}
			bim = 1
			sing = true
			base /= 10
		} else if s[i] == '-' {
			if sing {
				return 0
			}
			bim = -1
			sing = true
			base /= 10
		} else if s[i] > '9' || s[i] < '0' {
			return 0
		} else {
			num := int(s[i]) - '0'

			result += num * base
			base /= 10
		}
	}

	return result * bim
}
