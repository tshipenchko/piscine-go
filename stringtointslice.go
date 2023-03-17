package piscine

func StringToIntSlice(str string) []int {
	out := make([]int, 0, len(str))
	for _, c := range str {
		out = append(out, int(c))
	}
	return out
}
