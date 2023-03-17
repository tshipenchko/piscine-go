package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"piscine"
)

var Quad = map[string]func(x, y int){
	"quadA": piscine.QuadA,
	"quadB": piscine.QuadB,
	"quadC": piscine.QuadC,
	"quadD": piscine.QuadD,
	"quadE": piscine.QuadE,
}

func main() {
	filename := filepath.Base(os.Args[0])
	args := os.Args[1:]

	f, ok := Quad[filename]

	if !ok {
		fmt.Println("Error: quad", filename, "not found.")
		os.Exit(1)
	}

	if len(args) != 2 {
		fmt.Println("Provide x y")
		os.Exit(1)
	}

	x, err1 := strconv.Atoi(args[0])
	y, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Failed to parse (x, y):", x, y)
		os.Exit(1)
	}

	f(x, y)
}
