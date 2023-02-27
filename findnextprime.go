package piscine

func PrivateIsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i*i < nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	for i := nb; true; i++ {
		if PrivateIsPrime(i) {
			return i
		}
	}
	return -1
}
