package piscine

func IsNumeric(s string) bool {
	for _, value := range s {
		if '0' < value || value > '9' {
			return false
		}
	}
	return true
}
