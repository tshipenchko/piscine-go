package piscine

func AppendRange(min, max int) []int {
	var out []int

	for i := min; i < max; i++ {
		out = append(out, i)
	}

	return out
}
