package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, e := range a {
		counts[e]++
	}

	count := 0
	for _, v := range counts {
		if v%2 == 1 {
			count++
		}
	}

	if count == 0 {
		return -1
	}

	return count
}
