package piscine

func DescendAppendRange(max, min int) []int {
	out := []int{}
	if min > max {
		return out
	}
	for i := max; i > min; i-- {
		out = append(out, i)
	}
	return out
}
