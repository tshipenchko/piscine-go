package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	out := make(map[string]int)

	for _, e := range split(str, " ") {
		out[e]++
	}

	return out
}

func split(s, sep string) []string {
	IsCurrentSep := func(i int) bool {
		if len(s) >= len(sep)+i {
			return s[i:i+len(sep)] == sep
		}
		return false
	}

	var out []string
	var tmp []rune
	j := 0

	for i, char := range s {
		if i < j {
			continue
		}
		if IsCurrentSep(i) {
			j = i + len(sep)
			if len(tmp) != 0 {
				out = append(out, string(tmp))
				tmp = tmp[:0]
			}
		} else {
			tmp = append(tmp, char)
		}
	}

	if len(tmp) != 0 {
		out = append(out, string(tmp))
	}

	return out
}
