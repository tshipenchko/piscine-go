package piscine

func ReverseMenuIndex(menu []string) []string {
	out := make([]string, len(menu))

	for i, item := range menu {
		out[len(menu)-i-1] = item
	}

	return out
}

