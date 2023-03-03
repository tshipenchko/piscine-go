package piscine

func AppendRange(min, max int) []int {
	out := []int{}

	for i := min; i < max; i++ {
		out = append(out, i)
	}

	return out
}
