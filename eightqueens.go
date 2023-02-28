package piscine

import "github.com/01-edu/z01"

func PrintSolution(a []int) {
	for _, position := range a {
		z01.PrintRune(rune('1' + position)) // From 1, because position is index, but order is required
	}
	z01.PrintRune('\n')
}

func isNotInSlice(a []int, x int) bool {
	for _, value := range a {
		if value == x {
			return false
		}
	}
	return true
}

func newAddSlice(slice []int, x int) []int {
	new_slice := make([]int, len(slice)+1)
	copy(new_slice, slice)
	new_slice[len(slice)] = x // Add element to the end
	return new_slice
}

func EightQueensIteratorHelper(j, n, i int, a, b, c []int) {
	if j >= n {
		return
	}
	if isNotInSlice(a, j) && isNotInSlice(b, i+j) && isNotInSlice(c, i-j) {
		EightQueensIterator(
			n,
			i+1,
			newAddSlice(a, j),
			newAddSlice(b, i+j),
			newAddSlice(c, i-j),
		)
	}
	EightQueensIteratorHelper(j+1, n, i, a, b, c)
}

func EightQueensIterator(n, i int, a, b, c []int) {
	if i < n {
		EightQueensIteratorHelper(0, n, i, a, b, c)
	} else {
		PrintSolution(a)
	}
}

func EightQueens() {
	// Interpreted algorithm from https://en.wikipedia.org/wiki/Eight_queens_puzzle#Sample_program
	EightQueensIterator(
		8,
		0,
		make([]int, 0, 8),
		make([]int, 0, 8),
		make([]int, 0, 8),
	)
}
