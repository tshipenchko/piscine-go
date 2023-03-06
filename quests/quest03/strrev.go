package piscine

func StrRev(s string) string {
	dst := ""
	l := len(s)

	for i := 0; i < l; i++ {
		dst += string(s[l-i-1])
	}

	return dst
}
