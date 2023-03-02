package piscine

func Capitalize(s string) string {
	IsAlpha := func(r rune) bool {
		return 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || '0' <= r && r <= '9'
	}

	runes := []rune(s)

	for i, char := range runes {
		if i == 0 || !IsAlpha(runes[i-1]) && IsAlpha(char) {
			if 'a' <= char && char <= 'z' {
				runes[i] -= 32 // To upper
			}
		} else {
			if 'A' <= char && char <= 'Z' {
				runes[i] += 32 // To lower
			}
		}
	}

	return string(runes)
}
