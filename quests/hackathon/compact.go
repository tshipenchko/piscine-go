package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, e := range *ptr {
		if e != "" {
			(*ptr)[count] = e
			count++
		}
	}

	*ptr = (*ptr)[:count]
	return count
}
