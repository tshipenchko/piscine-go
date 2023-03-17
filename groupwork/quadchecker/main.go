package piscine

import (
	"fmt"
	"io"
	"os"
)

func CheckQuad() {
	stdin, _ := io.ReadAll(os.Stdin)
	input := string(stdin)

	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	if input[len(input)-1] != '\n' {
		input += "\n"
	}

	y := 0
	x := 0
	for i, v := range input {
		if v == '\n' {
			y++
		}
		if y == 0 {
			x = i + 1
		}

	}

	first := input[0]
	counter := 0
	if first == 'o' && input == QuadA(x, y) {
		PrintSol(x, y, &counter, 'A')
	}
	if first == '/' && input == QuadB(x, y) {
		PrintSol(x, y, &counter, 'B')
	}
	if first == 'A' && input == QuadC(x, y) {
		PrintSol(x, y, &counter, 'C')
	}
	if first == 'A' && input == QuadD(x, y) {
		PrintSol(x, y, &counter, 'D')
	}
	if first == 'A' && input == QuadE(x, y) {
		PrintSol(x, y, &counter, 'E')
	}

	if counter == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println()
	}
}

func PrintSol(x, y int, counter *int, l rune) {
	if *counter > 0 {
		fmt.Printf(" || ")
	}
	fmt.Printf("[quad%s] [%d] [%d]", string(l), x, y)
	*counter++
}
