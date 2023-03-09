package piscine

func Abort(a, b, c, d, e int) int {
	return BubbleSort([5]int{a, b, c, d, e})[3]
}

func BubbleSort(array [5]int) [5]int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
