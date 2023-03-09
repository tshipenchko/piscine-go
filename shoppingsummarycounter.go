package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	out := make(map[string]int)

	for _, e := range split(str) {
		out[e]++
	}

	return out
}

func split(s string) []string {
	out := []string{}

	j := 0
	for i, e := range s {
		if e == ' ' || i == len(s)-1 {
			out = append(out, s[j:i])
			j = i
		}
	}

	return out
}
