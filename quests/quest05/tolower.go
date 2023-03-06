package piscine

func ToLower(s string) string {
	result := ""

	for _, value := range s {
		if 'A' <= value && value <= 'Z' {
			result += string(value + 32)
		} else {
			result += string(value)
		}
	}

	return result
}
