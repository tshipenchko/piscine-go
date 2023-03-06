package piscine

import "github.com/01-edu/z01"

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] { // Ascending order
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	if n == 0 {
		z01.PrintRune('0')
	}

	digits := make([]int, 0, 10)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	BubbleSort(digits)

	for _, digit := range digits {
		z01.PrintRune(rune('0' + digit))
	}
}
