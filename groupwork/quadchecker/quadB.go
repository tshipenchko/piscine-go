package piscine

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 { // wrong argumets
		return ""
	}

	out := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // left up corner
				out += "/"
			} else if i == y && j == 1 { // left down corner
				out += "\\"
			} else if i == 1 && j == x { // right up  corner
				out += "\\"
			} else if i == y && j == x { // right down corner
				out += "/"
			} else if i == 1 || i == y { // horizontal line
				out += "*"
			} else if j == 1 || j == x { // vertical line
				out += "*"
			} else { // body
				out += " "
			}
		}
		out += "\n"
	}

	return out
}
