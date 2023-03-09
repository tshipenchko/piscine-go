package piscine

func Compact(ptr *[]string) int {
	out := make([]string, 0, len(*ptr))

	for _, e := range *ptr {
		if e != "" {
			out = append(out, e)
		}
	}

	*ptr = out
	return len(out)
}
