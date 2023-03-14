package main

import (
	"fmt"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				PrintString("ERROR: open ", filename, ": no such file or directory\n")
				os.Exit(1)
			}
			io.Copy(os.Stdout, file)
		}
	} else {
		fmt.Println("sadds")
		io.Copy(os.Stdout, os.Stdin)
	}
}

func PrintString(strings ...string) {
	for _, s := range strings {
		for _, char := range s {
			z01.PrintRune(char)
		}
	}
}
