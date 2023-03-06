package piscine

func IsAlpha(s string) bool {
	for _, value := range s {
		if !('A' <= value && value <= 'Z' || 'a' <= value && value <= 'z' || '0' <= value && value <= '9') {
			return false
		}
	}
	return true
}
