package main

import (
	"fmt"

	"piscine"
)

func TestD() {
	fmt.Println("D 5-3")
	piscine.QuadD(5, 3)
	fmt.Println()

	fmt.Println("D 5-1")
	piscine.QuadD(5, 1)
	fmt.Println()

	fmt.Println("D 1-1")
	piscine.QuadD(1, 1)
	fmt.Println()

	fmt.Println("D 1-5")
	piscine.QuadD(1, 5)
	fmt.Println()
}
