package piscine

func Rot14(s string) string {
	var out string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = rune((c-'a'+14)%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			c = rune((c-'A'+14)%26 + 'A')
		}
		out += string(c)
	}
	return out
}
