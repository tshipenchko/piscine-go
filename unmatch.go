package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, e := range a {
		counts[e]++
	}

	for k, v := range counts {
		if v%2 == 1 {
			return k
		}
	}

	return -1
}
