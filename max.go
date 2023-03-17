package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for _, e := range a {
		if e > max {
			max = e
		}
	}
	return max
}
