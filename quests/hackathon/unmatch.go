package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, e := range a {
		counts[e]++
	}

	for _, e := range a {
		if counts[e]%2 == 1 {
			return e
		}
	}

	return -1
}
