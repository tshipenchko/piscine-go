package piscine

func AppendRange(min, max int) []int {
	out := make([]int, 0)

	for i := min; i < max; i++ {
		out = append(out, i)
	}

	return out
}
