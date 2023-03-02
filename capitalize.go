package piscine

func Capitalize(s string) string {
	IsLetter := func(r rune) bool {
		return 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z'
	}

	runes := []rune(s)

	for i, char := range runes {
		if i == 0 || !IsLetter(runes[i-1]) && IsLetter(char) {
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
