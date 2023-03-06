package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}
