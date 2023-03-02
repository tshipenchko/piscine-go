package piscine

func Join(strs []string, sep string) string {
	result := ""

	for i, str := range strs {
		result += str
		if i+1 != len(strs) {
			result += sep
		}
	}

	return result
}
