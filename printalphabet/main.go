package main

import "github.com/01-edu/z01"

func main() {
	// fmt.Println("abcdefghijklmnopqrstuvwxyz")

	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(97 + i))
	}
	z01.PrintRune('\n')
}
