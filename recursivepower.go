package piscine

func RecursivePowerIterator(result, nb, power int) int {
	if power == 1 {
		return result * nb
	}
	return RecursivePowerIterator(result*nb, nb, power-1)
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return RecursivePowerIterator(1, nb, power)
}
