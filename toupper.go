package piscine

const RANGE = 'a' - 'A'

func ToUpper(s string) string {
	result := ""

	for _, value := range s {
		if 'a' <= value && value <= 'z' {
			result += string(value - RANGE)
		} else {
			result += string(value)
		}
	}

	return result
}
