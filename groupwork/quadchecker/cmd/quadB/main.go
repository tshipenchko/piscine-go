package main

import (
	"fmt"

	"piscine"
)

func main() {
	x, y, ok := piscine.GetNumber()
	if !ok {
		return
	}

	fmt.Print(piscine.QuadB(x, y))
}