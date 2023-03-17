package main

import (
	"fmt"

	"piscine"
)

func TestE() {
	fmt.Println("E 5-3")
	piscine.QuadE(5, 3)
	fmt.Println()

	fmt.Println("E 5-1")
	piscine.QuadE(5, 1)
	fmt.Println()

	fmt.Println("E 1-1")
	piscine.QuadE(1, 1)
	fmt.Println()

	fmt.Println("E 1-5")
	piscine.QuadE(1, 5)
	fmt.Println()
}
