package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 30 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	return RecursiveFactorial(nb - 1)
}
