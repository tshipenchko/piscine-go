package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	src := join(args)

	hasVovel := false
	for _, c := range src {
		if isVowel(c) {
			hasVovel = true
		}
	}

	if hasVovel {
		PrintString(mirrorVowelsInString(src))
	} else {
		PrintString(src)
	}

	z01.PrintRune('\n')
}

func mirrorVowelsInString(s string) string {
	vowelPositions := []int{}

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
		if IsUpper(runes[vowelPositions[i]]) {
			runes[vowelPositions[i]] = ToUpper(runes[vowelPositions[i]])
			runes[vowelPositions[j]] = ToUpper(runes[vowelPositions[j]])
		} else {
			runes[vowelPositions[i]] = ToLower(runes[vowelPositions[i]])
			runes[vowelPositions[j]] = ToLower(runes[vowelPositions[j]])
		}
		runes[vowelPositions[i]], runes[vowelPositions[j]] = runes[vowelPositions[j]], runes[vowelPositions[i]]
	}

	return string(runes)
}

func PrintString(s string) {
	for _, c := range s {
		z01.PrintRune(rune(c))
	}
}

func isVowel(c rune) bool {
	switch ToLower(c) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func join(ss []string) string {
	result := ""

	for i, str := range ss {
		result += str
		if i+1 != len(ss) {
			result += " "
		}
	}

	return result
}
