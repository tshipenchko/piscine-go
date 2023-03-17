package piscine

func Join(strs []string, sep string) string {
	out := ""
	for i, s := range strs {
		out += s
		if i != len(strs)-1 {
			out += sep
		}
	}
	return out
}
