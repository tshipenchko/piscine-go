package piscine

func MakeRange(min, max int) []int {
	var out []int

	if min >= max {
		return out
	}

	out = make([]int, max-min)
	for i := min; i < max; i++ {
		out[i-min] = i
	}

	return out
}
