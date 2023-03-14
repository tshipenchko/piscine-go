package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	out := 0
	for start != 1 {
		if out++; start%2 == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}
	}

	return out
}
