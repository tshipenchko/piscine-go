package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	count := int64(BasicAtoi(args[1]))
	hasMissingFiles := false

	filenames := args[2:]
	for i, filename := range filenames {
		file, err_open := os.Open(filename)
		stat, err_stat := os.Stat(filename)

		if err_open != nil || err_stat != nil {
			PrintString("open ", filename, ": no such file or directory\n")
			hasMissingFiles = true
			continue
		}

		if i != 0 {
			PrintString("\n")
		}

		if len(filenames) != 1 {
			PrintHeader(filename)
		}

		content := make([]byte, count)
		file.ReadAt(content, stat.Size()-count)
		PrintString(string(content))

		file.Close()
	}

	if hasMissingFiles { // TODO: fix exit status 1 message
		os.Exit(1)
	}
}

func PrintString(strings ...string) {
	for _, s := range strings {
		for _, char := range s {
			z01.PrintRune(char)
		}
	}
}

func PrintHeader(filename string) {
	PrintString("==> ", filename, " <==\n")
}

func BasicAtoi(s string) int {
	result := 0

	base := 1
	for i := 1; i < len(s); i++ {
		base *= 10
	}

	for i := 0; i < len(s); i++ {
		num := int(s[i]) - '0'
		result += num * base
		base /= 10
	}

	return result
}
