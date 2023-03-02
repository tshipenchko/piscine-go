package piscine

func IsUpper(s string) bool {
	for _, value := range s {
		if value < 'A' || value > 'Z' {
			return false
		}
	}
	return true
}
