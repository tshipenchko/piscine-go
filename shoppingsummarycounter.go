package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	out := make(map[string]int)

	for _, e := range split(str) {
		out[e]++
	}

	return out
}

func split(s string) []string {
	if len(s) == 0 {
		return []string{"", ""}
	}

	out := make([]string, 0, len(s))

	t := ""
	for _, e := range s {
		if e == ' ' {
			out = append(out, t)
			t = ""
		} else {
			t += string(e)
		}
	}
	out = append(out, t)

	return out
}
