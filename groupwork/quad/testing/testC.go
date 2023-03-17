package main

import (
	"fmt"

	"piscine"
)

func TestC() {
	fmt.Println("C 5-3")
	piscine.QuadC(5, 3)
	fmt.Println()

	fmt.Println("C 5-1")
	piscine.QuadC(5, 1)
	fmt.Println()

	fmt.Println("C 1-1")
	piscine.QuadC(1, 1)
	fmt.Println()

	fmt.Println("C 1-5")
	piscine.QuadC(1, 5)
	fmt.Println()
}
