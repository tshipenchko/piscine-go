package main

import "os"

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	a, ok_a := Atoi(args[0])
	b, ok_b := Atoi(args[2])

	if !ok_a || !ok_b {
		return
	}

	operator := args[1]
	out := 0

	if operator == "+" {
		if (out < a) != (b < 0) {
			return
		}
		out = a + b
	} else if operator == "-" {
		if (out < a) != (b > 0) {
			return
		}
		out = a - b
	} else if operator == "/" {
		if b == 0 {
			PrintString("No division by 0\n")
			return
		}

		out = a / b
	} else if operator == "*" {
		if a != 0 && (out/a != b) {
			return
		}
		out = a * b
	} else if operator == "%" {
		if b == 0 {
			PrintString("No modulo by 0\n")
			return
		}

		out = a % b
	} else {
		return
	}

	PrintNbr(out)
	PrintString("\n")
}
