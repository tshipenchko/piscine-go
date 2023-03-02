package piscine

func IsLower(s string) bool {
	for _, value := range s {
		if value < 'a' || value > 'z' {
			return false
		}
	}
	return true
}
