package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i < 26; i++ { // there is 26 letters in alphabet
		z01.PrintRune(rune('a' + i))
	}
	z01.PrintRune('\n')
}
