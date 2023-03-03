package main

import "github.com/01-edu/z01"

const RANGE = 'a' - 'A'

func IsUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func IsLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func ToUpper(r rune) rune {
	if IsLower(r) {
		return r - RANGE
	}
	return r
}

func ToLower(r rune) rune {
	if IsUpper(r) {
		return r + RANGE
	}
	return r
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
