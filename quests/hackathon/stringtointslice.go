package piscine

func StringToIntSlice(str string) []int {
	if len(str) == 0 {
		return nil
	}
	out := make([]int, 0, len(str))
	for _, c := range str {
		out = append(out, int(c))
	}
	return out
}
