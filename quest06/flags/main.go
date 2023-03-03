package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	flag_help := false
	flag_order := false
	flag_insert := ""
	src := ""

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			flag_help = true
		} else if arg == "-o" || arg == "--order" {
			flag_order = true
		} else if len(arg) >= 2 && arg[:2] == "-i" {
			flag_insert = arg[3:]
		} else if len(arg) >= 8 && arg[:8] == "--insert" {
			flag_insert = arg[9:]
		} else {
			src = arg
		}
	}

	// Main logic
	if flag_insert != "" {
		src += flag_insert
	}
	if flag_help || src == "" {
		fmt.Print(
			"--insert\n" +
				"  -i\n" +
				"	 This flag inserts the string into the string passed as argument.\n" +
				"--order\n" +
				"  -o\n" +
				"	 This flag will behave like a boolean, if it is called it will order the argument.\n",
		)
		return
	}
	if flag_order {
		src = string(BubbleSort([]rune(src)))
	}

	fmt.Println(src)
}

func BubbleSort(array []rune) []rune {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
