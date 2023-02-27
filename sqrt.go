package piscine

func SqrtIterator(nb, n int) int {
	if n*n == nb {
		return n
	}
	if n == nb/2 {
		return 0
	}
	return SqrtIterator(nb, n+1)
}

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	return SqrtIterator(nb, 1)
}
