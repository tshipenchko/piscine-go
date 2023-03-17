package main

import (
	"fmt"

	"piscine"
)

func TestB() {
	fmt.Println("B 5-3")
	piscine.QuadB(5, 3)
	fmt.Println()

	fmt.Println("B 5-1")
	piscine.QuadB(5, 1)
	fmt.Println()

	fmt.Println("B 1-1")
	piscine.QuadB(1, 1)
	fmt.Println()

	fmt.Println("B 1-5")
	piscine.QuadB(1, 5)
	fmt.Println()
}
