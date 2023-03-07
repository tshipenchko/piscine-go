package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for j := 1; j < len(a); j++ {
		i := j - 1
		if f(a[i], a[j]) > 0 {
			return false
		}
	}

	return true
}
