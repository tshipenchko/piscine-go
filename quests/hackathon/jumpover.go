package piscine

func JumpOver(str string) string {
	out := ""
	for i := 2; i < len(str); i += 3 {
		out += string(str[i])
	}
	out += "\n"
	return out
}
