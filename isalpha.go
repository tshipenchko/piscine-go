package piscine

func IsAlpha(s string) bool {
	for _, value := range s {
		if !('A' <= value && value <= 'Z' || 'a' <= value && value <= 'z') {
			return false
		}
	}
	return true
}
