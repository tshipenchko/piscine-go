package main

import "os"

func main() {
	args := os.Args[1:]
	src := join(args)

	PrintString(mirrorVowelsInString(src))
	PrintString("\n")
}

func mirrorVowelsInString(s string) string {
	vowelPositions := make([]int, 0, len(s))

	for i, r := range s {
		if isVowel(r) {
			vowelPositions = append(vowelPositions, i)
		}
	}

	if len(vowelPositions) == 0 {
		return s
	}

	runes := []rune(s)

	for i, j := 0, len(vowelPositions)-1; i < j; i, j = i+1, j-1 {
		runes[vowelPositions[i]], runes[vowelPositions[j]] = runes[vowelPositions[j]], runes[vowelPositions[i]]
	}

	return string(runes)
}
