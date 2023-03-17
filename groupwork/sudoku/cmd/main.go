package main

import (
	"os"
	"sudoku"
)

func main() {
	args := os.Args[1:]
	table, suc := sudoku.ParseSudoku(args)
	// TODO: rename suc to suck

	if !suc || !sudoku.ValidateSudoku(&table) { // smotrite ValidateSudoku
		sudoku.Println("Error")
	} else if sudoku.SolveSudoku(&table) {
		sudoku.PrintSudoku(table)
	} else {
		sudoku.Println("Error")
	}
}
