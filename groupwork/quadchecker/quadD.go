package piscine

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 { // wrong argumets
		return ""
	}

	/*
		ABBBC
		B   B
		ABBBC
	*/

	out := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // left up corner
				out += "A"
			} else if i == y && j == 1 { // left down corner
				out += "A"
			} else if i == 1 && j == x { // right up  corner
				out += "C"
			} else if i == y && j == x { // right down corner
				out += "C"
			} else if i == 1 || i == y { // horizontal line
				out += "B"
			} else if j == 1 || j == x { // vertical line
				out += "B"
			} else { // body
				out += " "
			}
		}
		out += "\n"
	}

	return out
}
