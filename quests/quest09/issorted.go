package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for j := 1; j < len(a); j++ {
		i := j - 1
		if f(a[i], a[j]) < 0 {
			ascending = false
			break
		}
	}

	for j := 1; j < len(a); j++ {
		i := j - 1
		if f(a[i], a[j]) > 0 {
			descending = false
			break
		}
	}

	return ascending || descending
}
