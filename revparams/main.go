package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := range os.Args[1:] {
		for _, sym := range os.Args[len(os.Args[1:])-i] {
			z01.PrintRune(rune(sym))
		}
		z01.PrintRune('\n')
	}
}
