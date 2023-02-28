package piscine

func AlphaCount(s string) int {
	count := 0

	for _, value := range s {
		if 'A' <= value && value <= 'Z' || 'a' <= value && value <= 'z' {
			count++
		}
	}

	return count
}
