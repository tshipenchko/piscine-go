package piscine

func FindNextPrime(nb int) int {
	for i := nb; true; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return -1
}
