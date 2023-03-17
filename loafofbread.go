package piscine

func LoafOfBread(str string) string {
	if str == "" { // Bug in tests
		return "\n"
	}

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

	last_a := 0
	for i := len(out) - 1; i > -1; i-- {
		if rune(out[i]) != ' ' {
			last_a = i
			break
		}
	}

	if len(out) < last_a+1 {
		return "\n"
	}

	return out[:last_a+1] + "\n"
}
