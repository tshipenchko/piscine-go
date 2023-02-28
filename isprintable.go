package piscine

func IsPrintable(s string) bool {
	for _, value := range s {
		if value < 32 || value > 127 {
			return false
		}
	}
	return true
}
