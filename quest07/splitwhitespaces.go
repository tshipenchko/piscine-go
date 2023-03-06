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

	var tmp []rune
	for _, char := range s {
		if !_IsSeparator(char) {
			tmp = append(tmp, char)
		} else if len(tmp) != 0 {
			out = append(out, string(tmp))
			tmp = tmp[:0]
		}
	}

	if len(tmp) != 0 {
		out = append(out, string(tmp))
	}

	return out
}
