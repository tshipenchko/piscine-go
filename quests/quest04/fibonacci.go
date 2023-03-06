package piscine

func FibonacciIterator(index, a, b int) int {
	if index == 1 {
		return b
	}
	return FibonacciIterator(index-1, b, a+b)
}

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return FibonacciIterator(index, 0, 1)
}
