package piscine

func ActiveBits(n int) int {
	c := 0

	for n > 0 {
		c += n % 2
		n /= 2
	}

	return c
}
