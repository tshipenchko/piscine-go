package piscine

import (
	"os"
	"strconv"
)

func GetNumber() (x, y int, ok bool) {
	args := os.Args[1:]

	if len(args) != 2 {
		return 0, 0, false
	}

	x, err1 := strconv.Atoi(args[0])
	y, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		return x, y, false
	}

	if x < 0 || y < 0 {
		return x, y, false
	}

	return x, y, true
}
