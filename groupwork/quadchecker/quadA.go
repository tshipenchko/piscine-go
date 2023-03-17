package piscine

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 { // wrong argumets
		return ""
	}

	out := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 || i == y) && (j == 1 || j == x) { // corners
				out += "o"
			} else if i == 1 || i == y { // horizontal
				out += "-"
			} else if j == 1 || j == x { // vertical
				out += "|"
			} else { // body
				out += " "
			}
		}
		out += "\n"
	}

	return out
}
