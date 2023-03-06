package piscine

func ConcatParams(args []string) string {
	out := ""

	for i, arg := range args {
		out += arg
		if i != len(args)-1 {
			out += "\n"
		}
	}

	return out
}
