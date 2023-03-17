package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	out := ""

	i := 0
	for _, e := range str {
		if i%6 != 5 && e == ' ' {
			continue
		}

		if i%6 == 5 {
			out += " "
		} else {
			out += string(e)
		}
		i++
	}

	return out + "\n"
}
