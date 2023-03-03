package main

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
