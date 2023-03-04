package piscine

func SplitWhiteSpaces(s string) []string {
	_IsSeparator := func(char rune) bool {
		switch char {
		case '\n', '\t', ' ':
			return true
		}
		return false
	}

	var out []string

	j := 0
	for i, char := range s {
		if _IsSeparator(char) || i == len(s)-1 {
			out = append(out, s[j:i])
			j = i + 1
		}
	}

	return out
}
