package main

import "github.com/01-edu/z01"

func main() {
	points := &point{}

	setPoint(points)

	PrintString("x = ")
	PrintInt(points.x)
	PrintString(", y = ")
	PrintInt(points.y)
	PrintString("\n")
}

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func PrintInt(i int) {
	out := ""
	for i > 0 {
		out = string(rune('0'+i%10)) + out
		i /= 10
	}
	PrintString(out)
}
