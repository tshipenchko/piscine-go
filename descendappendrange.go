package piscine

func DescendAppendRange(max, min int) []int {
	if min > max {
		return []int{}
	}
	out := make([]int, 0, max-min)
	for i := max; i > min; i-- {
		out = append(out, i)
	}
	return out
}
