package main

import (
	"fmt"

	"piscine"
)

func TestA() {
	fmt.Println("A 5-3")
	piscine.QuadA(5, 3)
	fmt.Println()

	fmt.Println("A 5-1")
	piscine.QuadA(5, 1)
	fmt.Println()

	fmt.Println("A 1-1")
	piscine.QuadA(1, 1)
	fmt.Println()

	fmt.Println("A 1-5")
	piscine.QuadA(1, 5)
	fmt.Println()
}
