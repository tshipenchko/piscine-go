package piscine

func Index(s string, toFind string) int {
	for i := range s {
		for j := 0; i+j < len(s); j++ {
			if j == len(toFind) {
				return i
			}
			if toFind[j] != s[i+j] {
				break
			}
		}
	}

	return -1
}
