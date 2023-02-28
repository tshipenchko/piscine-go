package piscine

import "github.com/01-edu/z01"

func PrintSolution(a [8]int) {
	for _, position := range a {
		if position == -420 {
			break
		}
		z01.PrintRune(rune(49 + position)) // From 1, because position is index, but order is required
	}
	z01.PrintRune(10)
}

func isNotInSlice(a [8]int, x int) bool {
	for _, value := range a {
		if value == x {
			return false
		}
		if value == -420 { // got to the end of slice
			return true
		}
	}
	return true
}

func MyAppend(a [8]int, x int) [8]int {
	for i := 0; i < len(a); i++ {
		if a[i] == -420 {
			a[i] = x
			break
		}
	}
	return a
}

func EightQueensIteratorHelper(j, n, i int, a, b, c [8]int) {
	if j >= n {
		return
	}
	if isNotInSlice(a, j) && isNotInSlice(b, i+j) && isNotInSlice(c, i-j) {
		EightQueensIterator(
			n,
			i+1,
			MyAppend(a, j),
			MyAppend(b, i+j),
			MyAppend(c, i-j),
		)
	}
	EightQueensIteratorHelper(j+1, n, i, a, b, c)
}

func EightQueensIterator(n, i int, a, b, c [8]int) {
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
		[8]int{-420, -420, -420, -420, -420, -420, -420, -420}, // -420 is special value that determines end of alice
		[8]int{-420, -420, -420, -420, -420, -420, -420, -420},
		[8]int{-420, -420, -420, -420, -420, -420, -420, -420},
	)
}
