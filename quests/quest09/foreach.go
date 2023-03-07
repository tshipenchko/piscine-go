package piscine

func ForEach(f func(int), a []int) {
	for _, x := range a {
		f(x)
	}
}
