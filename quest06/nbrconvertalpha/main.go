package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	args := make([]string, len(os.Args)-1)
	copy(args, os.Args[1:])

	flag_upper := false
	if args[0] == "--upper" {
		flag_upper = true
		args = args[1:]
	}

	for _, arg := range args {
		number := Atoi(arg)

		if number == -1 || number > 26 {
			z01.PrintRune(' ')
		} else if flag_upper {
			z01.PrintRune(rune('A' + number - 1))
		} else {
			z01.PrintRune(rune('a' + number - 1))
		}
	}

	z01.PrintRune('\n')
}

func Atoi(s string) int {
	result := 0

	base := 1
	for i := 1; i < len(s); i++ {
		base *= 10
	}

	for i := 0; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			return -1
		}

		num := int(s[i]) - '0'

		result += num * base
		base /= 10
	}

	return result
}
