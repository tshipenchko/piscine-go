package piscine

func BasicJoin(elems []string) string {
	result := ""

	for i, str := range elems {
		result += str
		if i+1 != len(elems) {
			result += ""
		}
	}

	return result
}
